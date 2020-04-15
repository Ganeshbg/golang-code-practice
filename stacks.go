package main

import (
	"fmt"
)

func main() {
	var stac []string
	stac = append(stac, "World!")
	stac = append(stac, "Hello")
	for len(stac) > 0 {
		n := len(stac) -1
		fmt.Println(stac[n])
		stac = stac[:n]
	}
}