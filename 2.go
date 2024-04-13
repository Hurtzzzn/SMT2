package main

import "fmt"

const NMax = 1000

type TAB [NMax]int

func jumlahBilangan(bilangan TAB, n int, jumlah *int) {
	var i int
	*jumlah = 0
	for i = 0; i < n; i++ {
		*jumlah += bilangan[i]
	}
}

func bacaData(bil *TAB, n *int) {
	var i, x int
	fmt.Scan(&x)
	i = 0
	for x >= 0 {
		bil[i] = x
		i++
		fmt.Scan(&x)
	}
	*n = i
}

func main() {
	var bilangan TAB
	var n, jumlah int
	bacaData(&bilangan, &n)
	jumlahBilangan(bilangan, n, &jumlah)
	fmt.Println(jumlah)
}
