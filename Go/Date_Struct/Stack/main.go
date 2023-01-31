package main

type stack struct {
	array []int
}

func (s *stack) Push(i int) {
	s.array = append(s.array, i)
}

func (s *stack) Pop() int {
	l := len(s.array) - 1
	toRemove := s.array[l]
	s.array = s.array[:l]
	return toRemove
}

func main() {

}
