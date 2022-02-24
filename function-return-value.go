package main

import "fmt"

func sayHello(name string) string {
	if name == "" {
		return "Hello there"
	} else {
		return "Hello" + name
	}
}
func main() {
	result := sayHello("Gontang")
	fmt.Println(result)
}
