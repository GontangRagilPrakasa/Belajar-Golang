package main

import "fmt"

func getFullName() (string, string) {
	return "Gontang", "Prakasa"
}
func main() {
	// firstname, lastname := getFullName()
	firstname, _ := getFullName()
	fmt.Println(firstname)
}
