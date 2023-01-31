package main

import (
	"fmt"
)

func main() {
	fixed := [...]int{1, 2, 3}
	fmt.Println(fixed)
	//fixed = [...]int{4, 5, 6, 7}
	a := []int{3, 4}
	fmt.Println(a)
	a = []int{6, 5, 3}

	a = append(a, 5, 6, 3)

	fmt.Println(cap(a), len(a))

	b := make([]int, 5)
	fmt.Println(b)

	a = a[0:6]
	fmt.Println(a)

}
