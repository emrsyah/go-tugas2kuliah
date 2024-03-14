package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(banyakBilanganGenap(N))

}

func banyakBilanganGenap(n int) int {
	/*  fungsi mengembalikan mengembalikan sebuah bilangan bulat yang menyatakan
	banyaknya bilangan genap dari 2 hingga N */
	genap := 0
	if n == 2 {
		return 1
	}
	if n%2 == 0 {
		genap = 1
	}
	return genap + banyakBilanganGenap(n-1)
}
