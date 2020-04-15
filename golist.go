package main

import (
    "fmt"
    "container/list"
)

func main() {
    var intList list.List
    intList.PushBack(2)
    intList.PushBack(3)
    intList.PushBack(4)
    for element := intList.Front(); element != nil; element = element.Next() {
        fmt.Println(element.Value.(int))
    }
}