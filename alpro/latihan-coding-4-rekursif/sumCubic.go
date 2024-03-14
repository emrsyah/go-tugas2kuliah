package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	sumCubic(n)
	fmt.Println(n)
}

func sumCubic(n int) int {
	/*{ I.S. Terdefinisi sebuah bilangan bulat positif n.
	F.S. Fungsi mengembalikan hasil perpangkatan 3 */
	if n == 1 {
		return 1
	} else {
		tmp := sumCubic(n - 1)
		fmt.Println(tmp)
		return n + tmp // gunakan baris ini untuk memproses perhitungan rekursif
	}

}
