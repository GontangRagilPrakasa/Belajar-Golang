package main

func main() {
	var ujian = 80
	var absensi = 80

	var lulusUJian bool = ujian > 80
	var lulusAbsensi bool = absensi > 80

	var lulus = lulusAbsensi && lulusUJian
	println(lulus)
}
