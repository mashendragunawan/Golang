package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// var nim int
	// var nama string
	// var kelas string

	// fmt.Print("Masukan Nim Anda  : ")
	// fmt.Scanf("%d", &nim)

	// fmt.Print("Masukan Nama Anda : ")
	// fmt.Scanf("%s", &nama)

	// fmt.Print("Masukan Kelas Anda: ")
	// fmt.Scanf("%s", &kelas)

	// fmt.Println("nim   :", nim)
	// fmt.Println("nama  :", nama)
	// fmt.Println("kelas :", kelas)

	// fmt.Println("mashe")
	db, err := sql.Open("mysql", "mashe:386412@tcp(127.0.0.1:3306)/hendra")

	if err != nil {
		//fmt.Print("koneksi Error")
		panic(err.Error())

	}

	defer db.Close()
	// fmt.Println("koneksi Berhasil")

	insert, err := db.Query("INSERT INTO user VALUES ('2016','mashe','386412')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("data berhasil di input")
}
