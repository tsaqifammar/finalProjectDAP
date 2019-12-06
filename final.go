// PROGRAM ProjekTUBES

package main

/* 		DATA KELOMPOK				*/
/*		Judul :   Aplikasi Ojek Online
		Anggota kelompok : Muhammad Tsaqif Ammar 1301194222
						   Alif Adwitiya Pratama 1301190465


						   Kelebihan Kode:
						   1. Memiliki fungsi kreatif berupa lokasi imajiner dan penghitugan jarak imajiner.
						   2. Tidak ada variable yang tidak jelas kami sebisa mungkin menggunakan nama variable yang menjelaskan fungsi variable tersebut.
						   3. Memiliki error handling untuk beberapa masukan.
						   4. Memiliki komen untuk menjelaskan kegunaan function.
						   5. Memiliki fungsi clear screen yang berasal dari pengetikan perintah cmd.

*/

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var clear map[string]func()

type driver struct {
	nama       string
	plat       string
	histori    transactionTab
	rataRating float64
}

type transaksi struct {
	tanggal string
	rating  int
	harga   int
	asal    string
	tujuan  string
	jarak   float64
}

type lokasi struct {
	nama string
	x    int
	y    int
}

type driverTab struct {
	tabel [1000]driver
	size  int
}

type transactionTab struct {
	tabel [1000]transaksi
	size  int
}

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear") //clear linux
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //clear windows
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //pengecekan
	if ok {
		value()
	} else {
		panic("") //eror handler
	}
}

func main() {
	var pilih int
	var tabelDriver driverTab

	// Pilihan awal (Insert - View - Exit)
	for keluar := false; !keluar; {
		CallClear()
		fmt.Println("MENU :")
		fmt.Println("1. Insert")
		fmt.Println("2. View")
		fmt.Println("3. Exit")
		fmt.Print("Pilih menu : ")
		fmt.Scanln(&pilih)

		switch pilih {
		// Memilih menu insert
		case 1:
			insert(&tabelDriver)
		// Memilih menu view
		case 2:
			view(&tabelDriver)

		case 3:
			keluar = true
		}

	}
	fmt.Println("Terimakasih sudah menggunakan!")
}

func insert(tabeldriver *driverTab) {
	var pilih int

	tab := *tabeldriver

	for keluar := false; !keluar; {
		CallClear()
		fmt.Println("INSERT : ")
		fmt.Println("1. Driver")
		fmt.Println("2. Transaksi perjalanan")
		fmt.Println("3. Back")
		fmt.Print("Pilih menu : ")
		fmt.Scanln(&pilih)
		switch pilih {
		// Menambahkan data driver baru
		case 1:
			var orang driver
			fmt.Print("Masukkan nama : ")
			fmt.Scanln(&orang.nama)
			fmt.Print("Masukkan plat : ")
			fmt.Scanln(&orang.plat)

			tab.tabel[tab.size] = orang
			tab.size++

		// Memasukkan data transaksi yang dilakukan seorang driver
		case 2:
			var name, date string
			var indeks int
			var asal, tujuan lokasi
			var cost, bintang int
			var transaction transaksi
			var jarak float64

			adaDiDataBase := false
			for adaDiDataBase == false {
				fmt.Print("Masukkan nama : ")
				fmt.Scanln(&name)
				// Mencari driver dengan nama 'name' di dalam tabeldriver
				ketemu := false
				for i := 0; i < tab.size && !ketemu; i++ {
					if tab.tabel[i].nama == name {
						indeks = i
						ketemu = true
						adaDiDataBase = true
					}
				}
				if adaDiDataBase == false {
					fmt.Println("{nama tidak ditemukan, ulangi input}") //eror handling
					fmt.Println(" ")
				}
			}
			// meminta lokasi asal dan tujuan, beserta harga dan jarak antar 2 lokasi tsb.
			fmt.Println("{Lokasi yang ada : A, B, C, D, E}")
			fmt.Print("Masukkan lokasi asal : ")
			for asal.nama != "A" && asal.nama != "B" && asal.nama != "C" && asal.nama != "D" && asal.nama != "E" { //eror handling jika dimasukan selain lokasi
				fmt.Scanln(&asal.nama)
				asal.nama = strings.ToUpper(asal.nama)
				if asal.nama != "A" && asal.nama != "B" && asal.nama != "C" && asal.nama != "D" && asal.nama != "E" {
					fmt.Println("{Maaf asal tidak ditemukan, ulangi input}")
					fmt.Println(" ")
					fmt.Println("{Lokasi yang ada : A, B, C, D, E}")
					fmt.Print("Masukan lokasi asal : ")
				}
			}

			fmt.Print("Masukkan lokasi tujuan : ")

			for tujuan.nama != "A" && tujuan.nama != "B" && tujuan.nama != "C" && tujuan.nama != "D" && tujuan.nama != "E" { //eror handling jika dimasukan selain lokasi
				fmt.Scanln(&tujuan.nama)
				tujuan.nama = strings.ToUpper(tujuan.nama)
				if tujuan.nama != "A" && tujuan.nama != "B" && tujuan.nama != "C" && tujuan.nama != "D" && tujuan.nama != "E" {
					fmt.Println("{Maaf asal tidak ditemukan, ulangi input}")
					fmt.Println(" ")
					fmt.Println("{Lokasi yang ada : A, B, C, D, E}")
					fmt.Print("Masukan lokasi asal : ")
				}
			}
			jarak = dist(asal, tujuan)
			cost = 2000 * int(jarak)
			fmt.Printf("Jarak = %.1f\n", jarak)
			fmt.Printf("Harga = %d\n", cost)

			// meminta rating/reputasi
			for inputBintangValid := false; !inputBintangValid; { //eror handling
				fmt.Print("Masukkan rating : ")
				fmt.Scanln(&bintang)
				if bintang >= 0 && bintang <= 5 {
					inputBintangValid = true
				} else {
					fmt.Println("{Masukkan tidak valid,diharapkan memberi masukan antara 1-5}")
				}
			}
			// mencatat waktu dan tanggal
			dt := time.Now()
			date = fmt.Sprintf("%s", dt.Format("01-02-2006 15:04:05"))

			transaction.tanggal = date
			transaction.rating = bintang
			transaction.harga = cost
			transaction.asal = asal.nama
			transaction.tujuan = tujuan.nama
			transaction.jarak = jarak

			lenHistori := tab.tabel[indeks].histori.size
			tab.tabel[indeks].histori.tabel[lenHistori] = transaction

			tab.tabel[indeks].histori.size++

			// Rata2 rating
			tab.tabel[indeks].rataRating = rata2Rating(tab.tabel[indeks])

		case 3:
			keluar = true
		}

	}
	*tabeldriver = tab
}

// Fungsi mencari jarak antar lokasi
func dist(asal, tujuan lokasi) float64 {
	var tabLokasi [5]lokasi
	var iAsal, iTujuan int
	// Nama dan koordinat (x,y) lokasi-lokasi
	tabLokasi[0] = lokasi{"A", 0, 0}
	tabLokasi[1] = lokasi{"B", 3, 4}
	tabLokasi[2] = lokasi{"C", 6, -1}
	tabLokasi[3] = lokasi{"D", 10, 20}
	tabLokasi[4] = lokasi{"E", -4, -5}

	// Mencari indeks berapa lokasi asal dan tujuan pada tabLokasi
	for i := 0; i < len(tabLokasi); i++ {
		if asal.nama == tabLokasi[i].nama {
			iAsal = i
		}
		if tujuan.nama == tabLokasi[i].nama {
			iTujuan = i
		}
	}

	// Return jarak antar lokasi(titik) = Sqrt((x1-x2)^2 + (y1-y2)^2)
	return math.Sqrt(math.Pow(float64(tabLokasi[iAsal].x-tabLokasi[iTujuan].x), 2.0) + math.Pow(float64(tabLokasi[iAsal].y-tabLokasi[iTujuan].y), 2.0))
}

func rata2Rating(pengemudi driver) float64 {
	sum := 0
	for i := 0; i < pengemudi.histori.size; i++ {
		sum += pengemudi.histori.tabel[i].rating
	}
	return float64(sum) / float64(pengemudi.histori.size)
}

func view(tabeldriver *driverTab) {
	var pilih int

	tab := *tabeldriver
	CallClear()
	for keluar := false; !keluar; {
		fmt.Println("INSERT : ")
		fmt.Println("1. Search")
		fmt.Println("2. Edit")
		fmt.Println("3. Sort")
		fmt.Println("4. Delete")
		fmt.Println("5. Back")
		fmt.Print("Pilih menu : ")
		fmt.Scanln(&pilih)
		switch pilih {
		// Melakukan Pencarian
		case 1:
			var name, plate string
			var indeks, pilih2 int

			fmt.Println("SEARCH")
			fmt.Println("1. Berdasarkan nama")
			fmt.Println("2. Berdasarkan plat")
			fmt.Print("Pilih menu : ")
			fmt.Scanln(&pilih2)
			switch pilih2 {
			// Pencarian berdasarkan nama dengan sequencial search
			case 1:
				fmt.Print("Masukkan nama : ")
				fmt.Scanln(&name)
				adaDiDataBase := false
				for i := 0; i < tab.size && !adaDiDataBase; i++ {
					if tab.tabel[i].nama == name {
						indeks = i
						adaDiDataBase = true
					}
				}
				if adaDiDataBase == false {
					CallClear()
					fmt.Println("{Tidak ditemukan nama yang dicari}") //eror handler
				} else {
					CallClear()
					fmt.Printf("Nama : %s\n", tab.tabel[indeks].nama)
					fmt.Printf("Plat : %s\n", tab.tabel[indeks].plat)
					fmt.Printf("Rating : %.1f\n", tab.tabel[indeks].rataRating)
					fmt.Println("Histori transaksi : ")
					for i := 0; i < tab.tabel[indeks].histori.size; i++ {
						fmt.Printf("%d. Tanggal : %s\n", i+1, tab.tabel[indeks].histori.tabel[i].tanggal)
						fmt.Printf("   Asal    : %s\n", tab.tabel[indeks].histori.tabel[i].asal)
						fmt.Printf("   Tujuan  : %s\n", tab.tabel[indeks].histori.tabel[i].tujuan)
						fmt.Printf("   Jarak   : %.1f\n", tab.tabel[indeks].histori.tabel[i].jarak)
						fmt.Printf("   Harga   : %d\n", tab.tabel[indeks].histori.tabel[i].harga)
						fmt.Printf("   Rating  : %d\n", tab.tabel[indeks].histori.tabel[i].rating)
					}
				}
			// Pencarian berdasarkan plat dengan binary search
			case 2:
				fmt.Print("Masukkan plat : ")
				fmt.Scanln(&plate)

				// Pengurutan sebelum melakukan binary search (insertion)
				for i := 1; i < tab.size; i++ {
					for j := i; j > 0; j-- {
						if tab.tabel[j-1].plat > tab.tabel[j].plat {
							tab.tabel[j-1], tab.tabel[j] = tab.tabel[j], tab.tabel[j-1]
						}
					}
				}

				L, R := 0, tab.size-1
				adaDiDataBase := false
				for L <= R && !adaDiDataBase {
					mid := L + (R-L)/2
					if tab.tabel[mid].plat > plate {
						R = mid - 1
					} else if tab.tabel[mid].plat < plate {
						L = mid + 1
					} else {
						adaDiDataBase = true
						indeks = mid
					}
				}
				if adaDiDataBase == false {
					CallClear()
					fmt.Println("Tidak ditemukan plat yang ingin dicari") //eror handler
				} else {
					CallClear()
					fmt.Printf("Nama : %s\n", tab.tabel[indeks].nama)
					fmt.Printf("Plat : %s\n", tab.tabel[indeks].plat)
					fmt.Printf("Rating : %.1f\n", tab.tabel[indeks].rataRating)
					fmt.Println("Histori transaksi : ")
					for i := 0; i < tab.tabel[indeks].histori.size; i++ {
						fmt.Printf("%d. Tanggal : %s\n", i+1, tab.tabel[indeks].histori.tabel[i].tanggal)
						fmt.Printf("   Asal    : %s\n", tab.tabel[indeks].histori.tabel[i].asal)
						fmt.Printf("   Tujuan  : %s\n", tab.tabel[indeks].histori.tabel[i].tujuan)
						fmt.Printf("   Jarak   : %.1f\n", tab.tabel[indeks].histori.tabel[i].jarak)
						fmt.Printf("   Harga   : %d\n", tab.tabel[indeks].histori.tabel[i].harga)
						fmt.Printf("   Rating  : %d\n", tab.tabel[indeks].histori.tabel[i].rating)
					}
				}
			default:
				fmt.Println("Tidak ada pilihan menu selain 1 dan 2")
			}
		// Melakukan edit pada data
		case 2:
			var name string
			var pilih2, indeks int

			fmt.Println("EDIT")
			adaDiDataBase := false
			for adaDiDataBase == false {
				fmt.Print("Masukkan nama : ")
				fmt.Scanln(&name)
				ketemu := false
				for i := 0; i < tab.size && !ketemu; i++ { //error handler
					if tab.tabel[i].nama == name {
						indeks = i
						ketemu = true
						adaDiDataBase = true
					}
				}
				if adaDiDataBase == false {
					fmt.Println("{Nama tidak ditemukan, masukan ulang nama}")
					fmt.Println("")
				}
			}
			fmt.Println("1. Nama")
			fmt.Println("2. Plat")
			fmt.Print("Pilih yang mau diedit : ")
			for inputValid := false; !inputValid; {
				fmt.Scanln(&pilih2)
				switch pilih2 {
				// Mengedit nama
				case 1:
					fmt.Print("Masukkan nama baru : ")
					fmt.Scanln(&tab.tabel[indeks].nama)
					inputValid = true
					CallClear()
				// Mengedit plat
				case 2:
					fmt.Print("Masukkan plat baru : ")
					fmt.Scanln(&tab.tabel[indeks].plat)
					inputValid = true
					CallClear()
				default:
					fmt.Println("{Tidak ada pilihan selain 1 dan 2}")
				}
			}
		// Melakukan pengurutan
		case 3:
			var pilih2, pilih3 int
			fmt.Println("SORT")
			fmt.Println("1. Berdasarkan Nama")
			fmt.Println("2. Berdasarkan Rating")
			fmt.Print("Pilih menu : ")
			for inputBenar := false; !inputBenar; {
				fmt.Scanln(&pilih2)
				if pilih2 < 0 || pilih2 > 2 {
					fmt.Println("{Input tidak valid, diharapkan mengisi antara 1-2}")
				} else {
					inputBenar = true
				}
			}
			fmt.Println("1. Pengurutan menaik")
			fmt.Println("2. Pengurutan menurun")
			fmt.Print("Pilih menu : ")
			for inputBenar := false; !inputBenar; {
				fmt.Scanln(&pilih3)
				if pilih3 < 0 || pilih3 > 2 {
					fmt.Println("{Input tidak valid, diharapkan mengisi antara 1-2}")
				} else {
					inputBenar = true
				}
			}
			CallClear()
			fmt.Println("Nama / Plat / Rating")
			tabTemp := tab
			// Memilih pengurutan berdasarkan nama (Selection Sort)
			if pilih2 == 1 {
				// Pengurutan menaik
				for i := 0; i < tabTemp.size; i++ {
					minIdx := i
					for j := i; j < tabTemp.size; j++ {
						if tabTemp.tabel[j].nama < tabTemp.tabel[minIdx].nama {
							minIdx = j
						}

					}

					tabTemp.tabel[minIdx], tabTemp.tabel[i] = tabTemp.tabel[i], tabTemp.tabel[minIdx]
				}

				// Pengurutan menurun
				if pilih3 == 2 {
					// Membalikkan array yang sudah diurut menaik menjadi menurun
					for i := 0; i < tabTemp.size/2; i++ {
						n := tabTemp.size - i - 1
						tabTemp.tabel[n], tabTemp.tabel[i] = tabTemp.tabel[i], tabTemp.tabel[n]
					}
				}
			} else if pilih2 == 2 {
				// Memilih pengurutan berdasarkan rating (Insertion Sort)
				// Pengurutan menaik
				for i := 1; i < tabTemp.size; i++ {
					Temp := tabTemp.tabel[i]
					j := i
					for j > 0 && tabTemp.tabel[j-1].rataRating > Temp.rataRating {
						tabTemp.tabel[j] = tabTemp.tabel[j-1]
						j--
					}
					tabTemp.tabel[j] = Temp
				}

				// Pengurutan menurun
				if pilih3 == 2 {
					// Membalikkan array yang sudah diurut menaik menjadi menurun
					for i := 0; i < tabTemp.size/2; i++ {
						n := tabTemp.size - i - 1
						tabTemp.tabel[n], tabTemp.tabel[i] = tabTemp.tabel[i], tabTemp.tabel[n]
					}
				}
			}
			// Print sorted array
			for i := 0; i < tabTemp.size; i++ {
				if tabTemp.tabel[i].nama != "-1" {
					fmt.Printf("%s / %s / %.1f\n", tabTemp.tabel[i].nama, tabTemp.tabel[i].plat, tabTemp.tabel[i].rataRating)
				}
			}
		// Melakukan penghapusan data / deleting data
		case 4:
			var name string

			fmt.Println("DELETE")
			adaDiDataBase := false
			for adaDiDataBase == false {
				fmt.Print("Masukkan nama yang data nya ingin dihapus : ")
				fmt.Scanln(&name)
				selesaiHapus := false
				for i := 0; i < tab.size && !selesaiHapus; i++ { //eror handling gak ada nama untuk di delete
					if tab.tabel[i].nama == name {
						tab.tabel[i].nama = "-1"
						tab.tabel[i].plat = "-1"
						tab.tabel[i].rataRating = -1
						selesaiHapus = true
						adaDiDataBase = true
					}
				}
				if adaDiDataBase == false {
					fmt.Println("{Nama tidak ditemukan, ulangi inputan}")
					fmt.Println(" ")
				}
			}
			CallClear()
		}

		if pilih == 5 {
			keluar = true
		} else if pilih < 0 || pilih > 5 {
			fmt.Println("{Input tidak valid, diharapkan mengisi antara 1-5}")
		}
	}

	*tabeldriver = tab
}
