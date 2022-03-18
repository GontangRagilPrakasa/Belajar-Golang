package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Eko")
	data.PushBack("Kurniawan")
	data.PushBack("Khannedy")
	data.PushFront("Budi")

	// fmt.Println(data.Front().Value)
	// fmt.Println(data.Back().Value)

	//jika mau ambil berdasarkan urutan memakai Next() selanjutnya dan Prev() untuk sebelumnya

	data.Front().Next().Next()
	data.Back().Prev().Prev()

	// tetapi hati hati akan ada data nill jika salah menempatkan fungsi
	// kemudian bagaimana caranya??

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	for element1 := data.Back(); element1 != nil; element1 = element1.Prev() {
		fmt.Println(element1.Value)
	}
}
