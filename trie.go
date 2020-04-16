package main

import "fmt"

const num = 26

type trie struct {
	letter [num]*trie
	isWord bool
}

type Node struct {
	root *trie
}

func getNode() *trie {
	node := &trie{}
	for i := 0; i < num; i++ {
		node.letter[i] = nil
	}
	return node
}
func (n *Node) insert(word string) {
	if n.root == nil {
		rt := getNode()
		n.root = rt
	}
	n.root.insert(word)
}

func (t *trie) insert(word string) {
	var temp *trie
	temp = t
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if temp.letter[index] == nil {
			temp.letter[index] = getNode()
		}
		temp = temp.letter[index]
	}
	temp.isWord = true
}

func (n *Node) search(word string) bool {
	temp := n.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if temp.letter[index] != nil {
			temp = temp.letter[index]
		} else {
			break
		}
	}
	return temp.isWord
}

func main() {
	var r *Node = &Node{}
	r.insert("hello")
	r.insert("hel")
	r.insert("ganesh")
	fmt.Println(r.search("hello"))
	fmt.Println(r.search("gan"))
	fmt.Println(r.search("sound"))
}
