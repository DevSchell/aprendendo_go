package main

import "fmt"

var name string = "Rodolfo"
var age int = 25
var systemVersion string = "1.0.0"

const (
	active   = iota
	inactive = iota
	pending  = iota
)

func main() {
	fmt.Println("System Profile starting...")

	fmt.Println("Profile 1")
	fmt.Printf("Name: %v \n", name)
	fmt.Printf("Age: %v \n", age)
	fmt.Printf("System Version: %v \n", systemVersion)
	fmt.Printf("Status: %v \n", active)
	fmt.Println("")

	name = "Roberto"
	age = 28
	systemVersion = "1.1.1"

	fmt.Println("Profile 2")
	fmt.Printf("Name: %v \n", name)
	fmt.Printf("Age: %v \n", age)
	fmt.Printf("System Version: %v \n", systemVersion)
	fmt.Printf("Status: %v \n", pending)
	fmt.Println("")

	name = "Roberta"
	age = 30
	systemVersion = "1.2.1"

	fmt.Println("Profile 3")
	fmt.Printf("Name: %v \n", name)
	fmt.Printf("Age: %v \n", age)
	fmt.Printf("System Version: %v \n", systemVersion)
	fmt.Printf("Status: %v \n", inactive)
	fmt.Println("")

}
