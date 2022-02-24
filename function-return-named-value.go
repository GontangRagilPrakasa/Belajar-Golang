package main

import "fmt"

func getFullNamed() (firstname, lastname string) {
	firstname = "Gontang"
	lastname = "Prakasa"

	return
}
func main() {
	// firstname, lastname := getFullName()
	firstname, _ := getFullNamed()
	fmt.Println(firstname)
}
