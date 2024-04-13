package main

import "fmt"

func reamur(c float64) float64 {
	return (4.0 / 5.0) * c
}
func fahrenheit(c float64) float64 {
	return (9.0/5.0)*c + 32
}
func kelvin(c float64) float64 {
	return c + 273
}
func main() {
	var c, r, f, k, cAwal, cAkhir, step, i float64
	fmt.Scan(&cAwal, &cAkhir, &step)
	fmt.Printf("%10s %10s %10s %10s\n", "Celcius", "Reamur", "Fahrenheit", "Kelvin")
	for i = cAwal; i <= cAkhir; i += step {
		c = float64(i)
		r = reamur(c)
		f = fahrenheit(c)
		k = kelvin(c)
		fmt.Printf("%10.2f %10.2f %10.2f %10.2f\n", c, r, f, k)
	}
}
