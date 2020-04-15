package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Println("Enter number")
	fmt.Scanln(&x)

	if ((x & 1) == 1) {
		fmt.Println("X is odd")
	} else {
		fmt.Println("X is even")
	}

}