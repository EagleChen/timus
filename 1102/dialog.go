package main

import (
	"fmt"
)

type Node struct {
	children map[byte]*Node
}

func (node *Node) isLeaf() bool {
	if _, ok := node.children[0]; ok {
		return true
	} else {
		return false
	}
}

func (node *Node) add(str string, index int, root *Node) {
	if index == len(str) {
		node.children[0] = root
		return
	}
	key := str[index]
	if _, ok := node.children[key]; !ok {
		node.children[key] = &Node{make(map[byte]*Node)}
	}
	node.children[key].add(str, index+1, root)
}

func buildTree() Node {
	root := Node{make(map[byte]*Node)}
	strs := []string{"in", "input", "one", "out", "output", "puton"}
	for _, str := range strs {
		root.add(str, 0, &root)
	}
	return root
}

func check(root *Node, str string) bool {
	current := root
	for _, v := range []byte(str) {
		if next, ok := current.children[v]; ok {
			current = next
		} else {
			if current.isLeaf() {
				if next, ok := root.children[v]; ok {
					current = next
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}
	if current.isLeaf() {
		return true
	} else {
		return false
	}
}

func main() {
	var num int
	var str string

	fmt.Scan(&num)

	root := buildTree()

	for ; num > 0; num-- {
		fmt.Scan(&str)

		if check(&root, str) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
