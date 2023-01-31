package main

import "log"

// func min[T interface {}] (a T, b T) T {
// 	return a
// }

// func min[T any](a T, b T) T {
// 	return a
// }

// func min[T ~float64](a T, b T) T {
// 	return a
// }

// func min[T mainTypes](a T, b T) T {
// 	return a
// }

func min[T mainTypes](a T, b T) T {
	if a < b {
		return a
	}
	return b
}

type mainTypes interface {
	~float64 | int | string
}

func main() {
	// log.Println(min(1, 2))
	type superFloat float64
	var s superFloat = 1.3
	log.Println(min(1, s))
}
