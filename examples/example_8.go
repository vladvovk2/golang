package main

import "fmt"

func main() {
	var a int = 4
	var b float64 = float64(a)
	var c uint = uint(b)
	fmt.Println(a, b, c)
}