package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(cetak_pola(n))
}

func cetak_pola(n int) int {
	if n == 1 || n == 2 || n == 3 {
		return 1
	} else {
		return cetak_pola(n-1) + cetak_pola(n-2) + cetak_pola(n-3)
	}
}
