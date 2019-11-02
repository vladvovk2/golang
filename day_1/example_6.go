package main

import (
	"fmt"
	"math"
)

func pow(a, b, lim float64) float64 {
	if v := math.Pow(a, b); v < lim {
		return v
	}
	fmt.Printf("%g >= %g\n", a, b)
	return lim
}
func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
