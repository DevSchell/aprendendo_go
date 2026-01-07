package main

import "fmt"

// Atribuindo variáveis package-level scope
var x int
var y string
var z bool

func main() {
	fmt.Printf("%v, %v, %v\n", x, y, z)
}
