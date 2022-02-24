package main

import "fmt"

func goodBye(name string) string {
	return "Good Bye " + name
}
func main() {
	sayGoodBye := goodBye

	result := sayGoodBye("Eko")
	fmt.Println(result)
}
