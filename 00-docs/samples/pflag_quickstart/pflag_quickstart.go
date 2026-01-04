package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

// 定义变量接受flag参数
var host string
var port int
var debug bool

func main() {
	// 声明flag后支持参数
	// --host xxx
	// --port 8080
	// --debug
	pflag.StringVarP(&host, "host", "h", "localhost", "host to connect to")
	pflag.IntVar(&port, "port", 8080, "port to listen on")
	pflag.BoolVar(&debug, "debug", false, "debug mode")

	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "pflag_quickstart")
		pflag.PrintDefaults()
	}

	// 解析参数
	if !pflag.Parsed() {
		pflag.Parse()
	}

	// 使用参数
	fmt.Printf("flag variables:\nhost: %s\nport: %d\ndebug: %v\n", host, port, debug)

}
