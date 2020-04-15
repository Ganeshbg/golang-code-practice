package main

import (
	"fmt"
	"time"
	"math/rand"
)


func main() {
	m := make(map[int]bool)
	bingo := [5][5]int{}
	count := 0
	i := 0
	j := 0
	for x:=0; x<25; x++ {
		m[x+1] = false
	}
	for {
		if (i < 5) {
			if (j == 5) {
				j = 0
				i = i + 1
			}
		}
		count = 0
		x := ((time.Now().Second() + rand.Intn(25)) % 25) + 1
		if (m[x] == false ) {
			m[x] = true
				if (j < 5) {
					bingo[i][j] = x
					j = j + 1
				}
		}
		for _,val := range m {
			if (val == false) {
				count = count + 1
			}
		}
		if count == 0 {
			break
		}
	}
	for l:=0; l<5; l++ {
		fmt.Println(bingo[l])
		fmt.Println()
	}
}