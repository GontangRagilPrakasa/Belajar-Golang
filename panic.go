package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")
	message := recover()
	if message != nil {
		fmt.Println("Error dengan pesan ", message)
	}

}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi berjalan")

}
func main() {
	runApp(false)
	fmt.Println("pesan anda?")
}
