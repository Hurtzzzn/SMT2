package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	cetak_rerata_master(N)
}

func cetak_rerata_master(N int) {
	var i int
	var jumlah int
	var rata float64
	i = 1
	jumlah = 0
	rata = 0
	cetak_rerata_recurse(N, i, jumlah, rata)
}

func cetak_rerata_recurse(N, i, jumlah int, rata float64) {
	if !(i < N) {
		jumlah += i
		rata = float64(jumlah) / float64(i)
		fmt.Println(jumlah, rata)
	} else {
		jumlah += i
		rata = float64(jumlah) / float64(i)
		fmt.Println(jumlah, rata)
		i++
		cetak_rerata_recurse(N, i, jumlah, rata)
	}
}
