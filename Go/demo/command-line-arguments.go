package main

import (
	"os"
	"fmt"
)

func main() {
	argsWidthProg := os.Args

	argsWidthoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWidthProg)
	fmt.Println(argsWidthoutProg)
	fmt.Println(arg)
}
