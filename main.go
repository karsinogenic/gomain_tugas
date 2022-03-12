package main

import (
	"fmt"

	modul "github.com/karsinogenic/gomodul_tugas"
	//modulasi "github.com/karsinogenic/gomodul_tugas"
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
	//var minyak = modul.Item{"minyak", 10, 30000, 0, 0}
	fmt.Println("Selamat datang mau belanja apa ? ")
	//modul.Coba1(12)
	modul.Kasir(12, 12)
	modul.Coba1(12)
	//minyak.Belanja1()
	// welcome := modul.Item
	// welcome.SelamatDatang("Budi")

	//fmt.Println("1.Minyak,stock:%d,harga:%d", minyak.stock, minyak.harga)
	//modulasi.Coba(30)
	//mbanyak, mtotal := 0
	//fmt.Println(mbanyak, mtotal)
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
