package main

import "fmt"

func main() {
	var c byte
	fmt.Scanf("%c", &c)
	fmt.Printf("%c", lowToUpper(c))
}

func lowToUpper(c byte) byte {
	tmp := c - 32
	return tmp
}
