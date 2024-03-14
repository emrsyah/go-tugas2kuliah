package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	balikDash(x)
}

func balikDash(x int) {
	if x == 0 {
		return
	}
	fmt.Print(x % 10)
	if x/10 != 0 {
		fmt.Print("-")
	}
	if x/10 != 0 {
		balikDash(x / 10)
	}
}
