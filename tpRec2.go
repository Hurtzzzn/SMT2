package main
import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)
	fmt.Println(gcd_recurse_master(A, B))
}

func gcd_recurse_master(a, b int) int {
	var c int
	return gcd_recurse(a, b, c)
}

func gcd_recurse(a, b, c int) int {
	if !(b > 0) {
		return a
	} else {
		c = a % b
		a = b
		b = c
		return gcd_recurse(a, b, c)
	}
}
