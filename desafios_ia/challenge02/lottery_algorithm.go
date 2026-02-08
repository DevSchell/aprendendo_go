package main

import "fmt"

var x int

func main() {

	for x = 0; x < 100; x++ {
		if x%3 == 0 && x%5 == 0 {
			fmt.Println("Lucky you!")
		} else if x%3 == 0 {
			fmt.Println("Lucky")
		} else if x%5 == 0 {
			fmt.Println(" you")
		}
	}

	switch x%2 == 0 {
	case true:
		fmt.Println("It is even!")
		break
	case false:
		fmt.Println("It is odd!")
		break
	}

}
