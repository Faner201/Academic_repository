package main

import (
	"fmt"
)

func main() {
	// var name type = expression
	var integer int = 1
	// auto initializing zero
	var integer_one int

	// var integer_two, integer_three = 1, 2;

	var integer_two, string = 1, "Aboba"

	// name := expression

	boolean := false

	fmt.Println(integer)
	fmt.Println(integer_one)
	fmt.Println(integer_two, string)
	fmt.Println(boolean)

	// Pointeres

	// x := 1
	// p := &x
	var x int = 1
	var p *int = &x

	fmt.Println(*p)

	// Type

	type fahrenheit int
	type celsius int

	var f fahrenheit = 32
	var c celsius = 0

	c = celsius((f - 32) * 5 / 9)

	fmt.Println(f, c)

	// const

	const (
		a = 1
		b = 2
		d = 3
	)
	fmt.Println(a, b, d)

	const (
		zero int = iota
		one
		two
		three
	)
	fmt.Println(zero, one, two, three)
}
