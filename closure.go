package main

import "fmt"

func main() {
	name := "Gontang"
	counter := 0

	increment := func() {
		name = "budi"
		fmt.Println("increment")
		counter++
	}

	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
