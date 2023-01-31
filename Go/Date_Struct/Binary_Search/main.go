package main

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	value int
	left  *node
	right *node
}

type bst struct {
	root *node
	len  int
}

func (n node) Outpur() string {
	return strconv.Itoa(n.value)
}

func (b bst) Outpur() string {
	sb := strings.Builder{}
	b.inOrderTravesal(&sb)
	return sb.String()
}

func (b bst) inOrderTravesal(sb *strings.Builder) {
	b.inOrderTravesalByNode(sb, b.root)
}

func (b *bst) add(value int) {
	b.root = b.addByNode(b.root, value)
	b.len++
}

func (b *bst) addByNode(root *node, value int) *node {
	if root == nil {
		return &node{value: value}
	}

	if root.value > value {
		root.left = b.addByNode(root.left, value)
	} else {
		root.right = b.addByNode(root.right, value)
	}

	return root
}

func (b bst) inOrderTravesalByNode(sb *strings.Builder, root *node) {
	if root == nil {
		return
	}
	b.inOrderTravesalByNode(sb, root.left)
	sb.WriteString(fmt.Sprintf("%s", root))
	b.inOrderTravesalByNode(sb, root.right)
}

func (b bst) search(value int) (*node, bool) {
	return b.searchByNode(b.root, value)
}

func (b bst) searchByNode(root *node, value int) (*node, bool) {
	if root == nil {
		return nil, false
	}

	if value == root.value {
		return root, true
	} else if value < root.value {
		return b.searchByNode(root.left, value)
	} else {
		return b.searchByNode(root.right, value)
	}
}

func (b *bst) remove(value int) {
	b.removeByNode(b.root, value)
}

func (b *bst) removeByNode(root *node, value int) *node {
	if root == nil {
		return root
	}

	if value > root.value {
		root.right = b.removeByNode(root.right, value)
	} else if value < root.value {
		root.left = b.removeByNode(root.left, value)
	} else {
		if root.left == nil {
			return root.right
		} else {
			temp := root.left
			for temp.right != nil {
				temp = temp.right
			}

			root.value = temp.value

			root.left = b.removeByNode(root.left, value)
		}
	}

	return root
}

func main() {

}
