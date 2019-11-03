package main

import "fmt"

type Hello float64
type World int64

func main() {
	var a Hello = 1.23
	var b World = 10
	fmt.Println(World(a) + b)
	fmt.Println(b.String())
}

func (c World) String() string {
	return fmt.Sprintf("%g C", c)
}
