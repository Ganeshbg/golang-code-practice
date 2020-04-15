package main
import (
	"fmt"
)

type Node struct {
	key int
	left *Node
	right *Node
}

type Tree struct {
	root *Node
}
func (t *Tree) insert(data int) error {
	if t.root == nil {
		fmt.Println("root is nil")
		t.root = &Node{data,nil,nil}
		return nil
	}
	return t.root.insert(data)
}

func (n *Node) insert(data int) error{
	if n == nil {
		fmt.Println("error cannot insert")
	} else if data <= n.key {
		if n.left == nil {
			n.left = &Node{data,nil,nil}
			return nil
		}
		return n.left.insert(data)
	} else {
		if n.right == nil {
			n.right = &Node{data,nil,nil}
			return nil
		}
		return n.right.insert(data)
	}
	return nil
}

func (n *Node) traverse() {
	if n == nil {
		return
	}
	n.left.traverse()
	fmt.Println(n.key)
	n.right.traverse()
}

func main() {
	var t *Tree = &Tree{}
	
	t.insert(4)
	t.insert(2)
	t.insert(6)
	t.insert(3)
	t.insert(5)
	t.insert(10)
	fmt.Println(t.root.key)
	t.root.traverse()
}