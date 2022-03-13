package main

import (
	"fmt"

	modul "github.com/karsinogenic/gomodul_tugas/v2"
)

// func coba(item [2]string) {
// 	fmt.Println(item)
// }

func main() {
	var choice int
	var troli_item = []string{"a", "Terigu"}
	var troli_bnyk = []int{2, 0}
	var minyak = modul.Item{"minyak", 10, 30000, 0, 0}
	var terigu = modul.Item{"terigu", 15, 15000, 0, 0}
	//var flag = true
	saldo, nama, ismember := modul.AnonymousData()

	//for flag==true{
ulang2:
	fmt.Printf("Selamat datang mau belanja apa ? ")

ulang3:
	fmt.Printf("1.Minyak,stock:%d,harga:%d||2.Terigu,stock:%d,harga:%d||3.Lihat Saldo\n", minyak.Stock, minyak.Harga, terigu.Stock, terigu.Harga)
	fmt.Print("Pilihan anda : ")
	fmt.Scanln(&choice)
	//modul.Coba1(12)
	switch choice {
	case 1:
		Stock, Belanja, Ntotal := minyak.Belanja1()
		minyak.Nbelanja = Belanja
		minyak.Ntotal = Ntotal
		minyak.Stock = Stock
		//if flag=
		troli_item[0] = "minyak"
		troli_bnyk[0] = minyak.Nbelanja
		//fmt.Println(troli_item, troli_bnyk)

		//flag:=1
	case 2:
		Stock, Belanja, Ntotal := terigu.Belanja1()
		terigu.Nbelanja = Belanja
		terigu.Ntotal = Ntotal
		terigu.Stock = Stock
		//if flag=
		troli_item[1] = "Terigu"
		troli_bnyk[1] = terigu.Nbelanja
		//fmt.Println(troli_item, troli_bnyk)
	case 3:
		fmt.Printf("Nama : %s,Saldo : %d,Status Member : %t", nama, saldo, ismember)
		goto ulang2
	default:
		fmt.Println("Input salah")
		goto ulang2
	}
	total := minyak.Ntotal + terigu.Ntotal
	ulang3 := modul.Troli(troli_item, troli_bnyk, total) //function 1 return
	if ulang3 == true {
		goto ulang3
	} else {
		//nothing
	}

	Diskon := func(IsMember bool) float64 { //anonymous func
		switch IsMember {
		case true:
			return 0.1
		case false:
			return 0
		default:
			return 0
		}
	}

	totaldiskon, sisa := modul.Kasir(Diskon(ismember), saldo, total)
	fmt.Printf("\nAnda belanja sebesar %d rupiah, sisa saldo anda kini %d rupiah\n", totaldiskon, sisa)
	fmt.Println("===============Terimakasih===============")
}

//}
