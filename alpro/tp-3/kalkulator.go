package main

import "fmt"

func main() {
	var opt int
	opt = 0
	for opt != 4 {
		menu()
		fmt.Print("Pilih (1/2/3/4)? ")
		fmt.Scan(&opt)
		if opt == 1 {
			hitungJumlah()
		} else if opt == 2 {
			hitungKali()
		} else if opt == 3 {
			hitungBagi()
		}
	}
}

// 8 + 7 + 8

func menu() {
	fmt.Println("-----------------------")
	fmt.Printf("%15s\n", "M E N U")
	fmt.Println("-----------------------")
	fmt.Println("1. Hitung Penjumlahan")
	fmt.Println("2. Hitung Perkalian")
	fmt.Println("3. Hitung Pembagian")
	fmt.Println("4. Exit")
	fmt.Println("-----------------------")
}

func hitungJumlah() {
	var a, b float64
	fmt.Print("Masukkan dua bilangan yang akan dijumlahkan: ")
	fmt.Scan(&a, &b)
	fmt.Println("Hasil penjumlahan:", a+b)
}

func hitungKali() {
	var a, b float64
	fmt.Print("Masukkan dua bilangan yang akan dikalikan: ")
	fmt.Scan(&a, &b)
	fmt.Println("Hasil perkalian:", a*b)
}

func hitungBagi() {
	var a, b float64
	fmt.Print("Masukkan dua bilangan yang akan dibagikan: ")
	fmt.Scan(&a, &b)
	fmt.Println("Hasil pembagian:", a/b)
}
