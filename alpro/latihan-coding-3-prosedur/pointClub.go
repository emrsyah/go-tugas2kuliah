// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main

import "fmt"

func main() {
	var golMasuk, golKemasukan int
	pointKlub(golMasuk, golKemasukan)
}

// Buatlah prosedur saja

// Buatlah prosedur saja
func pointKlub(golMasuk, golKemasukan int) {
	/*IS : Terdefinisi bilangan golMasuk dan golKemasukan sebanyak 38 kali
	  FS : Menampilkan jumlah point, jumlah gol memasukkan, jumlah gol kemasukan,
	  dan selisih gol dalam 1 musim.
	*/
	var lgm, lgk int
	golMasuk = 0
	golKemasukan = 0
	selisih := 0
	poin := 0
	for i := 1; i <= 38; i++ {
		fmt.Scan(&lgm, &lgk)
		golMasuk += lgm
		golKemasukan += lgk
		if lgm > lgk {
			// fmt.Println("menang")
			poin += 3
			selisih += lgm - lgk
		} else if lgm == lgk {
			// fmt.Println("seri")
			poin += 1
		} else {
			// fmt.Println("kalah")
			selisih -= lgk - lgm
		}
	}
	fmt.Println(poin, golMasuk, golKemasukan, selisih)
}
