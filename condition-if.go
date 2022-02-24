package main

import "fmt"

func main() {
	var name = "Eko"

	if name == "Eko" {
		fmt.Println("Hallo " + name)
	} else if name == "Joko" {
		fmt.Println("Hallo " + name)
	} else {
		fmt.Println("Not Found ")
	}

	var length = len(name)

	if length > 5 {
		fmt.Println("Not Identifier")
	} else {
		fmt.Println("Not Identifier ")
	}

	// atau car kedua

	if panjang := len(name); panjang > 5 {
		fmt.Println("Not Identifier")
	} else {
		fmt.Println("Not Identifier ")
	}
}
