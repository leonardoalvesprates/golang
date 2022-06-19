package main

import "fmt"

func main() {
	var myInt int = 16
	var val, ok = "yes", true
	var name string

	name = "Leo"

	fmt.Println("myInt is:", myInt)
	fmt.Println("myInt times two:", myInt*2)
	fmt.Println("val is:", val)
	fmt.Println("ok is:", ok)
	fmt.Println("name:", name)

	myInt2 := 32
	fmt.Println("myInt2 is:", myInt2)
}
