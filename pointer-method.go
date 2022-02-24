package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Marrioed() {
	man.Name = "Mr. " + man.Name
	// fmt.Println("(method)", man.Name) // jika tanpa pointer
}
func main() {
	eko := Man{"Eko"}
	eko.Marrioed()

	fmt.Println(eko.Name)
}
