/*package main

import (
	"fmt"
)

func main() {
	a := [7]int{1,23,34,44,55,77,99}
	length := len(a)-1
	low := 0
	search_item := 99
	var found int = 0
	for (low <= length) {
		mid := (low + length)/2
		fmt.Println(a[mid],mid)
		if (a[mid]==search_item){
			found = 1
			break
		}else if search_item > a[mid] {
			low = mid + 1
		}else {
			length = mid - 1
		}
	}
	if found == 0 {
		fmt.Println("Not Found")
	}else {
		fmt.Println("Found")
	}
}*/