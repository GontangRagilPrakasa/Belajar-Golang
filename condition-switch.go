package main

import "fmt"

func main() {
	name := "Eko"

	switch name {
	case "Eko":
		fmt.Println("Hallo " + name)
	case "Joko":
		fmt.Println("Hallo " + name)
	default:
		fmt.Println("Not Found ")
	}

	//cara kedua

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Terlalu Pajang")
	case false:
		fmt.Println("Connected")
	}

	//cara ketiga
	panjang := len(name)
	switch {
	case panjang > 10:
		fmt.Println("Lumayan Panjang ")
	case panjang > 5:
		fmt.Println("Cukup Panjang ")
	default:
		fmt.Println("Connected ")
	}
}
