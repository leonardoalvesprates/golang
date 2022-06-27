package main

import (
	"fmt"
	"motd/message"
)

func main() {
	message := message.Greeting("name1", "hello")
	fmt.Println(message)
}

// func greeting(name string, message string) string {
// 	return fmt.Sprintf("%s, %s", message, name)
// }
