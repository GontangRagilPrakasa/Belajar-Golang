package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	// Married       bool // jika ini ditambahkan maka variabel budi error karena sangat sensitif
}

// cara pertama
// func greeting(customer Customer, name string) {
// 	fmt.Println("Hallo ", name, " My Name is", customer.Name)
// }

// cara kedua
func (customer Customer) greeting(name string) {
	fmt.Println("Hallo ", name, " My Name is", customer.Name)
}

func main() {
	var person Customer
	person.Name = "Gontang"
	person.Address = "Yogyakarta"
	person.Age = 1
	fmt.Println(person)

	joko := Customer{
		Name:    "Joko",
		Address: "Sleman",
		Age:     1,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Jakarta", 1}
	fmt.Println(budi)

	// cara pertama
	// greeting(person, "Prakasa")

	// cara kedua
	person.greeting("Prakasa")
}
