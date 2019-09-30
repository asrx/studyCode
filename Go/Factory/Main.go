package main

import "./base"
import (
	_ "./cls1"
	_ "./cls2"
)

func main() {

	c1 := base.Create("Class1")
	c1.Do()

	c2 := base.Create("Class2")
	c2.Do()
}


