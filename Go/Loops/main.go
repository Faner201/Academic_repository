package main

import (
	"fmt"
)

func main() {
	// for
	for i, j := 0, 1; i < 10 && j < 10; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	// foreach

	s := "Hello World"

	for _, c := range s {
		if c == ' ' {
			break
		}
		fmt.Printf("%c", c)
	}

iForLoop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				break iForLoop
			}

			fmt.Println("%d%d", i, j)
		}
		fmt.Println()
	}
}
