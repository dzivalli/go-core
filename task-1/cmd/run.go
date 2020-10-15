package main

import (
	"fmt"
	"os"
	"strconv"

	"task-1/pkg/fibonacci"
)

func main() {
	arg, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(fibonacci.Calculate(arg))
}
