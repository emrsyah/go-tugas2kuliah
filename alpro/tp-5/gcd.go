package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(gcd(a, b, 1))
}

func gcd(a, b, i int) int {
	// Ini dia asumsinya itu a lebih besar dari b, tapi walaupun nggak tetep bakal ketuker juga di loop selanjutnya.
	//  Misal a = 5 dan b = 30. a mod b sendiri itu kan itu 5 alias a itu sendiri, jadi di loop selanjutnya bakal ketuker.
	//  jadi intinya fungsi si a itu sebagai yang di mod, dan b ini sebagai yang ngemod-nya. jadi kalo b berhasil nge mod a ya berarti b itu adalah gcd nya. Kalo nggak ya bakalan terus makin kecil ampe ke gcd.
	//  makanya kondisi berhenti itu ketika b = 0 karena itu ngasih tau kalo misalkan si a ini  (yang di loop sebelumnya adalah b) itu adalah kalo dipake nge mod jadinya 0.
	fmt.Println("iterasi ke-", i)
	fmt.Println("a =", a, "b =", b)
	if b != 0 {
		fmt.Println("a mod b =", a%b)
	} else {
		fmt.Println("a mod b = b sudah 0")
	}
	if b == 0 {
		return a
	}
	return gcd(b, a%b, i+1)
}
