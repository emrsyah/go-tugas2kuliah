//buatlah prosedur saja

func secMax(max, secondMax *int) {
	/* I.S. terdefinisi bilangan terbesar max, bilangan terbesar kedua secondMax.
	   dan inputan barisan bilangan bulat yang berakhir jika bilangan bulat sama dengan 0
	   F.S. max adalah bilangan bulat terbesar dan secondMax adalah bilangan bulat terbesar kedua */
	var ipt int
	fmt.Scan(&ipt)
	*max = ipt
	*secondMax = ipt
	fmt.Scan(&ipt)
	for ipt != 0 {
		if ipt > *max {
			tmp := *max
			*max = ipt
			*secondMax = tmp
		} else if ipt > *secondMax {
			*secondMax = ipt
		}
		fmt.Scan(&ipt)
	}
}
