package main

import "fmt"

type Filter func(string) string

// cara pertama
// func sayHello(name string, filter func(string) string) {
// 	nameFilter := filter(name)

// 	fmt.Println("Hello", nameFilter)
// }

//cara kedua
func sayHello(name string, filter Filter) {
	nameFilter := filter(name)

	fmt.Println("Hello", nameFilter)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHello("Gontang", spamFilter)
	sayHello("Anjing", spamFilter)
}
