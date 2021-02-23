package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/toolkits/pkg/runner"
)

var (
	vers *bool
	help *bool
	conf *string

	version = "No Version Provided"
)

func init() {
	vers = flag.Bool("v", false, "display the version.")
	help = flag.Bool("h", false, "print this help.")
	conf = flag.String("f", "", "specify configuration file.")
	flag.Parse()

	if *vers {
		fmt.Println("Version:", version)
		os.Exit(0)
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	runner.Init()
	fmt.Println("runner.cwd:", runner.Cwd)
	fmt.Println("runner.hostname:", runner.Hostname)
}

func main() {
	// if runtime.GOOS == "linux" || runtime.GOOS == "darwin" || runtime.GOOS == "windows" {
	// 	run()
	// } else {
	// 	fmt.Println("unsupport server operation:", runtime.GOOS)
	// 	os.Exit(0)
	// }
	run()
}

// // linux
// func reportStart() {
// 	go report.LoopReport()
// }

// func jobStart() {
// 	go timer.Heartbeat()
// }

// func monStart() {
// 	sys.Init(config.Config.Sys)
// 	stra.Init()

// 	funcs.BuildMappers()
// 	funcs.Collect()

// 	//插件采集
// 	plugins.Detect()

// 	//进程采集
// 	procs.Detect()

// 	//端口采集
// 	ports.Detect()

// 	//初始化缓存，用作保存COUNTER类型数据
// 	cache.Init()

// 	//日志采集
// 	go worker.UpdateConfigsLoop()
// 	go worker.PusherStart()
// 	go worker.Zeroize()
// }

// func parseConf() {
// 	if err := config.Parse(); err != nil {
// 		fmt.Println("cannot parse configuration file:", err)
// 		os.Exit(1)
// 	}
// }

// func endingProc() {
// 	c := make(chan os.Signal, 1)
// 	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
// 	select {
// 	case <-c:
// 		fmt.Printf("stop signal caught, stopping... pid=%d\n", os.Getpid())
// 	}

// 	logger.Close()
// 	http.Shutdown()
// 	fmt.Println("portal stopped successfully")
// }

// // windows
// // auto detect configuration file
// func aconf() {
// 	if *conf != "" && file.IsExist(*conf) {
// 		return
// 	}

// 	*conf = "etc/win-agent.local.yml"
// 	if file.IsExist(*conf) {
// 		return
// 	}

// 	*conf = "etc/win-agent.yml"
// 	if file.IsExist(*conf) {
// 		return
// 	}

// 	fmt.Println("no configuration file for collector")
// 	os.Exit(1)
// }

// func pconf() {
// 	if err := config.Parse(*conf); err != nil {
// 		fmt.Println("cannot parse configuration file:", err)
// 		os.Exit(1)
// 	}
// }

// func ending() {
// 	c := make(chan os.Signal, 1)
// 	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
// 	select {
// 	case <-c:
// 		fmt.Printf("stop signal caught, stopping... pid=%d\n", os.Getpid())
// 	}

// 	logger.Close()
// 	http.Shutdown()
// 	fmt.Println("sender stopped successfully")
// }

// func start() {
// 	runner.Init()
// 	fmt.Println("collector start, use configuration file:", *conf)
// 	fmt.Println("runner.cwd:", runner.Cwd)
// }

// func reportStart() {
// 	if err := report.GatherBase(); err != nil {
// 		fmt.Println("gatherBase fail: ", err)
// 		os.Exit(1)
// 	}

// 	go report.LoopReport()
// }

// func initWbem() {
// 	// This initialization prevents a memory leak on WMF 5+. See
// 	// https://github.com/prometheus-community/windows_exporter/issues/77 and
// 	// linked issues for details.
// 	// thanks prometheus windows exporter community for this issues. by yimeng
// 	logger.Debug("Initializing SWbemServices")

// 	s, err := wmi.InitializeSWbemServices(wmi.DefaultClient)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	wmi.DefaultClient.AllowMissingFields = true
// 	wmi.DefaultClient.SWbemServicesClient = s
// }
