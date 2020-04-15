package main

import (
    "fmt"
)

func main() {
    var x, y int
    fmt.Println("Enter 2 numbers")
    fmt.Scanln(&x)
    fmt.Scanln(&y)
    if(x^y < 0) {
        fmt.Println("Opposite signs")
    } else {
        fmt.Println("Same signs")
    }

}