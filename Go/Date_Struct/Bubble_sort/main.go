package main

import (
	"log"
)

func swap(numbers []int, i int, j int) {
	temp := numbers[i]
	numbers[i] = numbers[j]
	numbers[j] = temp
}

func BubleSort(numbers []int) {
	for j := 0; j < len(numbers); j++ {
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[i+1] {
				swap(numbers, i, i+1)
			}
		}
	}
}

func main() {
	numbers := []int{6, 2, 4, 5, 5, 6, 2, 3}
	BubleSort(numbers)
	log.Println(numbers)
}
