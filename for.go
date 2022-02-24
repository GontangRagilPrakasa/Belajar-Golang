package main

import "fmt"

func main() {
	counter := 1

	for counter < 10 {
		// fmt.Println("Perulangan ke = ", counter)
		counter++
	}

	for counter1 := 1; counter1 <= 10; counter1++ {
		// fmt.Println("Perulangan ke = ", counter1)
	}

	slice := []string{"Gonrang", "Ragil", "Prakasa"}
	for i := 0; i < len(slice); i++ {
		fmt.Println("Perulangan ke = ", slice[i])
	}

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	// untuk megetahui var tidak dibutuhkan
	for _, value := range slice {
		fmt.Println("Index =", value)
	}

	person := make(map[string]string)
	person["name"] = "Eko"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
