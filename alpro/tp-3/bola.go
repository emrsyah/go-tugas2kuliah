package main

import "fmt"

// banyak pertandingan, jum kemenangan, jum draw, jum kalah, jum gol, jum kegolan, jum selisih gol, total poin
var jm, jd int = 0, 0

func main() {
	var n, g, k, jklh, jg, jkg, jsg, jp int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&g, &k)
		hitungMenang(g, k, &jm)
		hitungDraw(g, k, &jd)
		hitungKalah(g, k, &jklh)
		hitungJumGolKegolanSelisih(g, k, &jg, &jkg, &jsg)
		hitungJumPoint(&jp)
	}
	fmt.Println(n, jm, jd, jklh, jg, jkg, jsg, jp)
}

func hitungMenang(g, k int, jm *int) {
	if g > k {
		*jm += 1
	}
}

func hitungDraw(g, k int, jd *int) {
	if g == k {
		*jd += 1
	}
}

func hitungKalah(g, k int, jk *int) {
	if g < k {
		*jk += 1
	}
}

func hitungJumGolKegolanSelisih(g, k int, jg, jkg, jsg *int) {
	*jg += g
	*jkg += k
	*jsg += g - k
}

func hitungJumPoint(jp *int) {
	*jp = jm*3 + jd*1
}
