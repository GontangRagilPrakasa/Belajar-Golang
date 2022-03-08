package main

import (
	"belajar-golang/database"
	"fmt"
	// ini blank identifier menggunakan _ sebelum import _ "belajar-golang/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
