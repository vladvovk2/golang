package main

import "fmt"

func main() {
	x := 2
	p := &x

	fmt.Println(*p)
	*p += 3

	fmt.Println(x)
}
