package main

import "fmt"

func main() {
	ages := map[string]int{}
	ages["name1"] = 16

	// if
	if ages["name1"] < 18 {
		fmt.Println("Name1 can't vote yet")
	} else if ages["name1"] < 67 {
		fmt.Println("name1 is not of retirement age just yet")
	} else {
		fmt.Println("Name1 is of retirement age")
	}

	// switch
	switch {
	case ages["name1"] < 18:
		fmt.Println("Name1 can't vote yet")
	case ages["name1"] < 67:
		fmt.Println("name1 is not of retirement age just yet")
	default:
		fmt.Println("Name1 is of retirement age")
	}

	// switch
	switch ages["name1"] {
	case 1, 2, 3, 5, 7, 11, 13, 17, 19:
		fmt.Println("Name1 has a prime number age")
	case 16:
		fmt.Println("Name1 can drive")
	case 18:
		fmt.Println("Name1 can vote")
	case 67:
		fmt.Println("name1 can retire")
	default:
		fmt.Println("Name1's age is nothing special")
	}
}
