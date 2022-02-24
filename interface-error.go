package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagian int) (int, error) {
	if pembagian == 0 {
		return 0, errors.New("Pembagian tidak boleh 0")
	} else {
		result := nilai / pembagian
		return result, nil
	}
}
func main() {
	var contohError error = errors.New("ups error")
	fmt.Println(contohError.Error())

	hasil, err := Pembagian(100, 10)
	if err == nil {
		fmt.Println("Hasil ", hasil)
	} else {
		fmt.Println("error", err.Error())
	}
}
