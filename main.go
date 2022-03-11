package main

import (
	"fmt"

	modulasi "github.com/karsinogenic/gomodul_tugas"
)

func main() {
	fmt.Println("coba")
	bayar, pilihan := modulasi.Troli(100, 2)
	if pilihan == true {
		fmt.Println("belanja lagi")
	}
	fmt.Println("troli : ")
	fmt.Println("kembalinya :", modulasi.Kasir(10000, bayar))

}
