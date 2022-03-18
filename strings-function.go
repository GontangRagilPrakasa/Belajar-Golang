package main

import (
	"fmt"
	"strings"
)

func main() {
	// untuk mengecek apakah paramter Gontang ada dalam kalimat Gontang Ragil Prakasa
	// keluarannya adalah true and false
	fmt.Println(strings.Contains("Gontang Ragil Prakasa", "Gontang"))

	// untuk memotong menjadi string
	fmt.Println(strings.Split("Gontang Ragil Prakasa", " "))

	fmt.Println(strings.ToLower("Gontang Ragil Prakasa"))
	fmt.Println(strings.ToUpper("Gontang Ragil Prakasa"))
	fmt.Println(strings.ToTitle("gontang ragil prakasa"))
	fmt.Println(strings.Trim("     gontang ragil prakasa     ", " "))
	fmt.Println(strings.ReplaceAll("gontang gontang ragil prakasa", "gontang", "budi"))
}
