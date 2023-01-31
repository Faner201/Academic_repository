package main

type queue struct {
	array []int
}

func (q *queue) Enqueue(i int) {
	q.array = append(q.array, i)
}

func (q *queue) Remove() int {
	toRemove := q.array[0]
	q.array = q.array[1:]
	return toRemove
}

func main() {

}
