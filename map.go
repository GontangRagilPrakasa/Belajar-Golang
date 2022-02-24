package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Gontang Ragil Prakasa",
		"address": "Yogyakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// Cara kedua
	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Coding"
	book["author"] = "Gontang"
	book["ups"] = "Salah"

	delete(book, "ups")

	fmt.Println(book)
}
