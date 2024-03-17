package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	cetak_rerata(n, 0, 1)
}

func cetak_rerata(n, total, i int) {
	total += i
	rata := float64(total) / float64(i)
	if i == n {
		fmt.Println(total, rata)
	} else {
		fmt.Println(total, rata)
		cetak_rerata(n, total, i+1)
	}
}
