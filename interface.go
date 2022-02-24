package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasname HasName) {
	fmt.Println("Hello", hasname.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

// Cara kedua
type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var eko Person
	eko.Name = "Eko"
	sayHello(eko)

	cat := Animal{
		Name: "Cat",
	}
	sayHello(cat)
}
