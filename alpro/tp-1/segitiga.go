package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	for i := 1; i <= x; i++ {
		for j := 1; j <= x; j++ {
			if j != i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}
