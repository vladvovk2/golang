package main

import "fmt"

func main() {
	a := "Hello!"
	for _, a := range a {
		a := a + 'a' - 'a'
		fmt.Println(a)
	}
}
