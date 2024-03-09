package main

import "fmt"

func main() {
	var r, s, ll, lp, kl, kp, tl, tp float64
	fmt.Scan(&r, &s)
	if r == 0 && s == 0 {
		return
	}
	fmt.Printf("%7s %7s %7s %7s %7s %7s %7s %7s\n", "R", "S", "LL", "LP", "KL", "KP", "TL", "TP")
	for r != 0 || s != 0 {
		hitungLuasKelilingLingkaran(r, &ll, &kl)
		hitungLuasKelilingPersegi(s, &lp, &kp)
		hitungTotal(ll, lp, kl, kp, &tl, &tp)
		// fmt.Println(ll, lp, kl, kp)
		fmt.Printf("%7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f\n", r, s, ll, lp, kl, kp, tl, tp)
		fmt.Scan(&r, &s)
	}
}

func hitungLuasKelilingLingkaran(r float64, l *float64, k *float64) {
	/* I.S. r adalah radius dari lingkaran
	   F.S. l adalah luas lingkaran dengan radius r, dan k adalah keliling lingkaran dengan radius r */
	*l = 3.14 * r * r
	*k = 2 * 3.14 * r
}

func hitungLuasKelilingPersegi(s float64, l *float64, k *float64) {
	/* I.S. s adalah panjang sisi persegi
	   F.S. l adalah luas persegi panjang dengan sisi sisi, dan k adalah keliling persegi dengan sisi sisi */
	*l = s * s
	*k = 4 * s
}

func hitungTotal(ll, lp, kl, kp float64, toLuas *float64, toKeliling *float64) {
	// I.S. terdefinisi bilangan bilangan real ll luas lingkaran, lp luas persegi, kl keliling lingkaran, kp keliling persegi
	// F.S. toLuas dan toKeliling sebagai bilangan real hasil penjumlahan masing-masing luas ataupun keliling
	*toLuas = ll + lp
	*toKeliling = kl + kp
}
