package main

import "log"

const arraySize = 100

type HashTable struct {
	array [arraySize]*Bucket
}

type Bucket struct {
	head *BucketNode
}

type BucketNode struct {
	key  string
	next *BucketNode
}

func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % arraySize
}

func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &Bucket{}
	}
	return result
}

func (b *Bucket) insert(k string) {
	if !b.search(k) {
		newNode := &BucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		log.Println(k, "alredy exists")
	}
}

func (b *Bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (b *Bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

func main() {
	hashTable := Init()

	list := []string{
		"Lopata",
		"Nikola",
		"Polka",
	}

	for _, v := range list {
		hashTable.Insert(v)
	}

	log.Println(hashTable.Search("Lopat"))

}
