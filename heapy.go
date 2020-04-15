package main

import (
	"fmt"
	"container/heap"
)

type intheap []int

func (iheap intheap) Len() int {
	return len(iheap)
}

func (iheap intheap) Less(i,j int) bool {
	return iheap[i]<iheap[j]
}

func (iheap *intheap) Push(x interface{}) {
	*iheap = append(*iheap,x.(int))
}

func (iheap intheap) Swap(i,j int) { iheap[i], iheap[j] = iheap[j], iheap[i]}

func (iheap *intheap) Pop() interface{}{
	var last intheap = *iheap
	var n int = len(last)
	var x int
	x = last[n-1]
	*iheap = last[:n-1]
	return x
}

func main() {
	var inth *intheap = &intheap{3,6,1}
	heap.Init(inth)
	heap.Push(inth,0)
	for len(*inth) > 0 {
		fmt.Println("%d", heap.Pop(inth))
	}
}