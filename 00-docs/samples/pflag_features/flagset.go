package main

import "github.com/spf13/pflag"

var host string
var port int

var username string
var password string

func main() {
	// 父命令
	pflag.StringVarP(&host, "host", "h", "localhost", "host to connect to")
	pflag.IntVarP(&port, "port", "p", 8080, "port to connect to")

	// 解析
	pflag.Parse()

}
