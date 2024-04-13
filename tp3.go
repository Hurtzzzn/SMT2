package main

import (
	"fmt"
	"math"
)

func panjangX(r, t float64) float64 {
	return r * math.Cos(t*math.Pi/180)
}
func panjangY(r, t float64) float64 {
	return r * math.Sin(t*math.Pi/180)
}
func main() {
	var r, t, x, y float64
	fmt.Scan(&r, &t)
	x = panjangX(r, t)
	y = panjangY(r, t)
	fmt.Printf("%.2f %.2f\n", x, y)
}
