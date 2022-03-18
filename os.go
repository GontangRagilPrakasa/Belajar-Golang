package main

import (
	"fmt"
	"os"
)

func main() {
	// More Info https://pkg.go.dev/os

	// Data Argument
	args := os.Args
	fmt.Println("Argument ")
	fmt.Println(args)

	fmt.Println(args[1])

	// Data Hostname atau

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname : ", hostname)
	} else {
		fmt.Println("Error : ", err.Error())
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSOWRD")
	// how to set export APP_USERNAME=root

	fmt.Println(username)
	fmt.Println(password)

}
