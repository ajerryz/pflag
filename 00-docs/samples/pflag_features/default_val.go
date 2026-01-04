package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

func main() {
	// bool 如果不传递是 true
	success := pflag.Bool("success", false, "should this flag be set")
	failure := pflag.Bool("failure", true, "should this flag be set")

	pflag.Parse()

	fmt.Printf("Success: %v\n", *success)
	fmt.Printf("Failure: %v\n", *failure)
}
