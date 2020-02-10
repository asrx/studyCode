package main

import (
	"os"
	"strings"
	"fmt"
)

func main() {
	for _,e := range os.Environ(){
		pair := strings.Split(e,"=")
		fmt.Println(pair[0])
	}
}
