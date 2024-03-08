// Buatlah prosedur saja

func cetakHitungJumlahRata(n, m int, jum *int, rata *float64) {
	/* I.S.: Terdefinisi bilangan bulat n dan m
	   F.S.: jum berisi nilai penjumlahan bilangan bulat dari n hingga m
	         dan rata berisi nilai rata-rata bilangan bulat dari n hingga m */
	*jum = 0
	for i := n; i <= m; i++ {
		fmt.Println(i)
		*jum += i
	}
	*rata = float64(*jum) / (float64(m) - float64(n) + 1.0)
}