package main

import (
	"fmt"
)

func main() {

	var num1 int
	var num2 int

	fmt.Print("Masukan Angka Pertama : ")
	fmt.Scanf("%d", &num1)

	fmt.Print("Masukan Angka Kedua :")
	fmt.Scanf("%d", &num2)

	pl := num1 / num2

	fmt.Println(num1, ":", num2, "=", pl)

	fmt.Println("Perhitungan Selesai sampai di sini Ok")

	// var nim int
	// var nama string

	// fmt.Print("Masukan Nama Nim : ")
	// fmt.Scanf("%d", &nim)

	// fmt.Print("Maukan Nama Anda : ")
	// fmt.Scanf("%s", &nama)

	// fmt.Println("==========================")

	// fmt.Println("Nim :", nim)
	// fmt.Println("Nama :", nama)
}
