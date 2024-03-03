// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main

import "fmt"

func main() {
	var aw, ak, b float64
	fmt.Scan(&aw, &ak, &b)
	fmt.Printf("%10s %10s %10s %10s \n", "Celcius", "Reamur", "Kelvin", "Fahrenheit")
	for i := aw; i <= ak; i += b {
		r := rea(i)
		f := fah(i)
		k := kel(i)
		//   fmt.Println(i,r,f,k)
		fmt.Printf("%10.2f %10.2f %10.2f %10.2f \n", i, r, f, k)
	}
}

func rea(c float64) float64 {
	return (c * 0.8)
}

func fah(c float64) float64 {
	return (c * 9.0 / 5.0) + 32
}

func kel(c float64) float64 {
	return c + 273.0
}

// fmt.Printf("%10.2s %10.2s %10.2s %10.2s \n", C, R, F, K)
