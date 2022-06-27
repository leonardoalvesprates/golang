package main

import "fmt"

func main() {
	ages := map[string]int{}
	ages["name1"] = 18
	ages["name2"] = 28
	ages["name3"] = 67
	ages["name4"] = 16
	ages["name5"] = 5

	// sequence looping
	for key, age := range ages {
		switch age {
		case 1, 2, 3, 5, 7, 11, 13, 17, 19:
			fmt.Println(key, "has a prime number age")
		case 16:
			fmt.Println(key, "can drive")
		case 18:
			fmt.Println(key, "can vote")
		case 67:
			fmt.Println(key, "can retire")
		default:
			fmt.Println(fmt.Sprintf("There's nothing special about %s's current age.", key))
		}
	}

	// traditional 'for' looping
	for i := 1; i <= 10; i++ {
		fmt.Println("We are counting:", i)
	}

	// conditional looping
	a := 0
	for a < 10 {
		fmt.Println("We're counting:", a)
		a++ // if absent this it is turns in a infinite loop
	}

	// stopping iterations
	b := 0
	for b < 10 {
		if b%2 == 0 {
			b++
			continue
		} else if b == 5 {
			break
		}

		fmt.Println("We're counting:", b)
		b++
	}
}
