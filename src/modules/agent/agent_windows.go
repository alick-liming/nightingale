// +build windows

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/didi/nightingale/src/modules/agent/cache"
	"github.com/didi/nightingale/src/modules/agent/config"
	"github.com/didi/nightingale/src/modules/agent/sys/identity"

	"github.com/didi/nightingale/src/modules/agent/http/routes"
	"github.com/didi/nightingale/src/modules/agent/report"
	"github.com/didi/nightingale/src/modules/agent/stra"
	"github.com/didi/nightingale/src/modules/agent/sys"
	"github.com/didi/nightingale/src/modules/agent/sys/funcs"
	"github.com/didi/nightingale/src/modules/agent/sys/plugins"
	"github.com/didi/nightingale/src/modules/agent/sys/ports"
	"github.com/didi/nightingale/src/modules/agent/sys/procs"

	tlogger "github.com/didi/nightingale/src/common/loggeri"
	"github.com/didi/nightingale/src/toolkits/http"

	"github.com/StackExchange/wmi"
	"github.com/gin-gonic/gin"
	"github.com/toolkits/pkg/file"
	"github.com/toolkits/pkg/logger"
	"github.com/toolkits/pkg/runner"
)

// var (
// 	vers *bool
// 	help *bool
// 	conf *string
// )

// func init() {
// 	vers = flag.Bool("v", false, "display the version.")
// 	help = flag.Bool("h", false, "print this help.")
// 	conf = flag.String("f", "", "specify configuration file.")
// 	flag.Parse()

// 	if *vers {
// 		fmt.Println("version:", config.Version)
// 		os.Exit(0)
// 	}

// 	if *help {
// 		flag.Usage()
// 		os.Exit(0)
// 	}
// }

func run() {
	aconf()
	pconf()
	start()

	initWbem()

	cfg := config.Get()

	tlogger.Init(cfg.Logger)

	identity.Init(cfg.Identity, cfg.IP)
	log.Println("endpoint & ip:", identity.GetIdent(), identity.GetIP())

	sys.Init(cfg.Sys)
	stra.Init(cfg.Stra)

	funcs.InitRpcClients()
	funcs.BuildMappers()
	funcs.Collect()

	//插件采集
	plugins.Detect()

	//进程采集
	procs.Detect()

	//端口采集
	ports.Detect()

	//初始化缓存，用作保存COUNTER类型数据
	cache.Init()
	if cfg.Enable.Report {
		reportStart()
	}

	r := gin.New()
	routes.Config(r)
	http.Start(r, "agent", cfg.Logger.Level)
	ending()
}

// auto detect configuration file
func aconf() {
	if *conf != "" && file.IsExist(*conf) {
		return
	}

	*conf = "etc/win-agent.local.yml"
	if file.IsExist(*conf) {
		return
	}

	*conf = "etc/win-agent.yml"
	if file.IsExist(*conf) {
		return
	}

	fmt.Println("no configuration file for collector")
	os.Exit(1)
}

// parse configuration file
func pconf() {
	if err := config.Parse(*conf); err != nil {
		fmt.Println("cannot parse configuration file:", err)
		os.Exit(1)
	}
}

func ending() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case <-c:
		fmt.Printf("stop signal caught, stopping... pid=%d\n", os.Getpid())
	}

	logger.Close()
	http.Shutdown()
	fmt.Println("sender stopped successfully")
}

func start() {
	runner.Init()
	fmt.Println("collector start, use configuration file:", *conf)
	fmt.Println("runner.cwd:", runner.Cwd)
}

func reportStart() {
	if err := report.GatherBase(); err != nil {
		fmt.Println("gatherBase fail: ", err)
		os.Exit(1)
	}

	go report.LoopReport()
}

func initWbem() {
	// This initialization prevents a memory leak on WMF 5+. See
	// https://github.com/prometheus-community/windows_exporter/issues/77 and
	// linked issues for details.
	// thanks prometheus windows exporter community for this issues. by yimeng
	logger.Debug("Initializing SWbemServices")

	s, err := wmi.InitializeSWbemServices(wmi.DefaultClient)
	if err != nil {
		log.Fatal(err)
	}
	wmi.DefaultClient.AllowMissingFields = true
	wmi.DefaultClient.SWbemServicesClient = s
}
