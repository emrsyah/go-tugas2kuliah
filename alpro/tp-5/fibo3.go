package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fibo3(n))
}

func fibo3(n int) int {
	if n == 1 || n == 2 || n == 3 {
		return 1
	} else {
		return fibo3(n-1) + fibo3(n-2) + fibo3(n-3)
	}
}
