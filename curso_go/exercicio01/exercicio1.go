package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	//1. Uma declaração Print
	fmt.Printf("%v, %v, %v\n", x, y, z)

	//2. Múltiplas declarações Print
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v", z)
}
