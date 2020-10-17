package main

import (
	"flag"
	"fmt"
	"task-1/pkg/fib"
)

func main() {
	var nFlag = flag.Int("n", 1, "fibonacci number")
	flag.Parse()
	fmt.Println(fib.Num(*nFlag))
}
