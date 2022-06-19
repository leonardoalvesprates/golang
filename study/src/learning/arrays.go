package main

import "fmt"

func main() {
	// oneline
	names := [3]string{"name1", "name2", "name3"}
	fmt.Println(names)
	fmt.Println("names:", names)

	// otherway
	var othernames [3]string
	othernames[0] = "name1"
	othernames[1] = "name2"
	// othernames[2] = "name3"
	// othernames[3] = "name4" // invalid argument: array index 3 out of bounds [0:3]
	fmt.Println(othernames)
	fmt.Println("othernames[2] is nil:", othernames[2] == "")
	fmt.Println("othernames[0]:", othernames[0])

	// not limited array - slices
	nolimited := []string{}
	nolimited = append(nolimited, "name1")
	nolimited = append(nolimited, "name2", "name3")
	fmt.Println("nolimited:", nolimited)

	// make function
	names2 := make([]string, 2)
	names2[0] = "name1"
	names2[1] = "name2"
	// names2[3] = "name3" // panic: runtime error: index out of range [3] with length 2
	fmt.Println(names2)
}
