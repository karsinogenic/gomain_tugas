package main

import (
	"fmt"
)

type item struct {
	nama     string
	stock    int
	harga    int
	nBelanja int
	nTotal   int
}

// func coba(item [2]string) {
// 	fmt.Println(item)
// }

func main() {
	var minyak = item{"minyak", 10, 30000, 0, 0}
	fmt.Println("Selamat datang mau belanja apa ? ")
	fmt.Println("1.Minyak,stock:%s,harga:%s", minyak.stock, minyak.harga)
	mbanyak, mtotal := minyak.Belanja1()
	fmt.Println(mbanyak, mtotal)
	//var terigu = item{"terigu", 15, 15000, 0}
	//var item = [2]string{"minyak", "terigu"}

	// coba(item)
	// fmt.Println(minyak, terigu)

	//banyak(item)

	// switch belanja {
	// case 1:
	// 	bayar, pilihan := modulasi.Troli(100, 2)
	// }
	// fmt.Println("coba")
	// bayar, pilihan := modulasi.Troli(100, 2)
	// if pilihan == true {
	// 	fmt.Println("belanja lagi")
	// }
	// fmt.Println("troli : ")
	// fmt.Println("kembalinya :", modulasi.Kasir(10000, bayar))

}
