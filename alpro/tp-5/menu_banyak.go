package main

import "fmt"

func main() {
	// fmt.Println(sumBilangan(5))
	// fmt.Println(fakBilangan(5))
	fmt.Println(fibo(6))
}

func sumBilangan(n int) int {
	if n == 1 {
		return 1
	} else {
		return n + sumBilangan(n-1)
	}
}

func fakBilangan(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * fakBilangan(n-1)
	}
}

func fibo(n int) int {
	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}
