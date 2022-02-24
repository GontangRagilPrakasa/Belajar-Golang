package main

import "fmt"

func sumAll(number ...int) int {
	total := 0

	for _, value := range number {
		total += value
	}

	return total
}
func main() {
	total := sumAll()
	fmt.Println(total)

	slice := []int{10, 20}
	total = sumAll(slice...)
	fmt.Println(total)
}
