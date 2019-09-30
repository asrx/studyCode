package main

import (
	"os"
	_"fmt"
	_"sort"
)

// Panic

func main() {

	panic("a problem")	
	
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}

}

/*

*/
