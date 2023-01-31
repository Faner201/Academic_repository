package main

import (
	"fmt"
)

func main() {
	var a [3]int
	fmt.Println(a)
	fmt.Println(a[0])

	a[0] = 5

	aCopy := a
	fmt.Println(aCopy == a)

	b := [4]int{5, 3, 4, 2}
	fmt.Println(b)

	c := [...]int{2, 3}
	fmt.Println(c)

	for _, v := range b {
		fmt.Println(v)
	}

	array := [...]int{0: 1}
	fmt.Println(array)
	array_one := [...]int{0: 1, 4: 5}
	fmt.Println(array_one)

	array2d := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(array2d)
	array3d := [2][2][2]int{array2d, array2d}
	fmt.Println(array3d)
}
