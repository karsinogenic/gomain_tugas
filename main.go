package main

import (
	"fmt"

	modulasi "github.com/karsinogenic/gomodul_tugas"
)

func main() {
	fmt.Println("coba")
	modulasi.Troli(10000, 2)
	fmt.Println("troli : ")
	fmt.Println("kembalinya :", modulasi.Kasir(10000, 2000))

}
