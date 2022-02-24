package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)

	// Pass by Reference dengan Pointer dalam artian parents dan child akan berubah

	address3 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address4 := &address3

	address4.City = "Bandung"

	fmt.Println(address3)
	fmt.Println(address4)

	// cara lain dari pointer memakai * ( ex: *Address )

	var address5 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address6 *Address = &address5

	address6.City = "Bandung"

	// address6 = &Address{"Malang", "Jawa Timur", "Indonesia"}	// Pointer dengan build ulang

	*address6 = Address{"Malang", "Jawa Timur", "Indonesia"} //operasi ini siapapun yang mengacu pada address utama akan ikut berubah
	fmt.Println(address5)
	fmt.Println(address6)

	// bagaimana kalau dengan pointer kosong
	var address7 *Address = new(Address)
	address7.City = "Ponorogo"
	fmt.Println(address7)

	// bagaimana perubahan paramter dengan pointer?
	var alamat = Address{
		City:     "Jogja",
		Province: "DIY",
		Country:  "",
	}
	// ChangeCountryToIndonesia(alamat) // tidak akan beruah
	ChangeCountryToIndonesia(&alamat) // akan berubah
	fmt.Println(alamat)

}
