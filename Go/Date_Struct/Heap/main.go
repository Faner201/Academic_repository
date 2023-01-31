package main

import "log"

type heap struct {
	array []int
}

func (h *heap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *heap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *heap) Extract() int {
	exctracted := h.array[0]
	l := len(h.array) - 1

	if h.array == nil {
		log.Println("cannot extract because array length is 0")
		return -1
	}
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)

	return exctracted
}

func (h *heap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	var childToCompare int

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *heap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func main() {
	m := &heap{}

	buildHeap := []int{10, 30, 40, 2, 52, 34, 123, 4, 643}
	for _, v := range buildHeap {
		m.Insert(v)
		log.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		log.Println(m)
	}
}
