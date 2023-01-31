package main

import "log"

func main() {
	//make() only works with slices, maps and channels but new() works with any

	// _ := make(int)

	s := make([]int, 0)
	i := new(int)
	log.Println(s, i)

	//Returned types
	//make() returns the target type but new() returns the address
	//Memory Initialization
	//make() initializes memory and new() zeroes memory
	var v map[int]int = make(map[int]int)
	var p *map[int]int = new(map[int]int)

	log.Println(v, p)
}
