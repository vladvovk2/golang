package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 1 { return sqrt(-x) + "i" }
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(9), sqrt(-4))
}