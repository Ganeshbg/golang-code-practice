package main

import (
	"fmt"
)

func mergeSort(list []int) []int {
	if len(list) == 1 {
		return list
	}
	var left,right []int
	length := len(list)
	mid := int(length/2)
	left = list[:mid]
	right = list[mid:length]
	return merge(mergeSort(left),mergeSort(right))
}

func merge(left, right []int) []int {
	result := make([]int,len(left) + len(right))
	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right [0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for len(left) > 0 {
		result[i] = left[0]
		left = left[1:]
		i++
	}

	for len(right) > 0 {
		result[i] = right[0]
		right = right[1:]
		i++
	}
	return result
}

func main() {
	x := []int{7,3,4,9,1,23,0}
	fmt.Println(x)
	y := mergeSort(x)
	fmt.Println(y)
}