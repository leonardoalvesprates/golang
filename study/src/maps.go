package main

import "fmt"

func main() {
	birthdays := map[string]string{
		"name1": "02/06/1990",
		"name2": "01/01/1957",
		"name3": "06/24/2975",
	}

	fmt.Println(birthdays)

	ages := map[string]int{}
	ages["name1"] = 28
	ages["name2"] = 61
	ages["name3"] = 43

	fmt.Println(ages, ages["name1"])

	// deleting a map
	delete(birthdays, "name1")
	fmt.Println(birthdays)

	//
	fmt.Print("name ---  birthday --- age \n")
	fmt.Println("name2:", birthdays["name2"], ages["name2"])

}
