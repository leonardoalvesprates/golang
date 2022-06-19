package main

import "fmt"

/*
this is a comment
on multiple lines
*/

// main is the primaty function
func main() {
	fmt.Println("Simple String") // comment here
	fmt.Println(`
This is a multi line
String, that can also contain "quotes".
`)
	fmt.Println("😉")
	fmt.Println("\u2272")
	fmt.Println('L')
	fmt.Println("\n")
	fmt.Println("Addition:", 1+3, "\n")
	fmt.Println("Substration:", 1-3, "\n")
	fmt.Println("Multiplication:", 1*3, "\n")
	fmt.Println("Division:", 20/3)
	fmt.Println("Division:", 20.0/3)
	fmt.Println("Great than 1 > 2:", 1 > 2)
	fmt.Println("Equivalent 4.0 = 4:", 4.0 == 4)
}
