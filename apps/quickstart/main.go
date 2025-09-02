package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

func main() {
	// 定义字符串参数,返回字符串指针
	name := pflag.StringP("name", "n", "default", "your name")
	// 定义整型参数，返回int指针
	age := pflag.IntP("age", "a", 0, "your age")
	// 定义布尔参数
	verbose := pflag.BoolP("verbose", "v", false, "verbose output")
	// 定义字符串切片
	hobbies := pflag.StringSliceP("hobby", "h", []string{}, "your hobbies")

	// 设置参数必须传递

	if !pflag.Parsed() {
		pflag.Parse() // 解析命令行参数
	}

	fmt.Printf("name:%v\n", *name)
	fmt.Printf("age:%v\n", *age)
	fmt.Printf("hobbies:%v\n", *hobbies)
	fmt.Printf("verbose:%v\n", *verbose)

	// 访问非 选项(flag)参数
	//pflag.Arg(1) // 获取第i个非 option参数
	fmt.Printf("args:%v\n", pflag.Args())
}
