package main

import (
	"fmt"
	"strings"
)

type node struct {
	value int
	next  *node
}

type linkedList struct {
	head *node
	len  int
}

func (n node) String() string {
	return fmt.Sprintf("%d", n.value)
}

func (l *linkedList) add(value int) {
	newNode := new(node)
	newNode.value = value

	if l.head == nil {
		l.head = newNode
	} else {
		iterator := l.head
		for ; iterator.next != nil; iterator = iterator.next {
		}
		iterator.next = newNode
	}
}

func (l *linkedList) remove(value int) {
	var previos *node

	for iterator := l.head; iterator != nil; iterator = iterator.next {
		if iterator.value == value {
			if l.head == iterator {
				l.head = iterator.next
			} else {
				previos.next = iterator.next
				return
			}
		}
	}
}

func (l linkedList) get(value int) *node {
	for iterator := l.head; iterator != nil; iterator = iterator.next {
		if iterator.value == value {
			return iterator
		}
	}
	return nil
}

func (l linkedList) Outpur() string {
	output := strings.Builder{}

	for iterator := l.head; iterator != nil; iterator = iterator.next {
		output.WriteString(fmt.Sprintf("%s", iterator))
	}
	return output.String()
}

func main() {
}
