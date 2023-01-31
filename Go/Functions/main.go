package main

import (
	"fmt"
)

func main() {
	varidic(1, 2, 6, 7, 34, 6)

	fmt.Println(fibonachi(10))

	i := 1
	defer second()
	withPoin(&i)
	fmt.Println(i)

	fmt.Println(returnFunction(1, returnParam))

	first()

	//Anonymuos function

	aFunction := func() {
		fmt.Println("Hello there")
	}

	func() {
		fmt.Println("Hello there 2")
	}()

	aFunction()

	aFunction = first
}

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

func returnParam(i int) int {
	return i
}

func returnFunction(i int, f func(int) int) int {
	return f(i)
}

func varidic(integers ...int) {
	for _, integer := range integers {
		fmt.Println(integer)
	}
}

func nameLenght(name string) (string, int) {
	return name, len(name)
}

func greating(name string) (string, bool) {
	if len(name) == 0 {
		return "", false
	} else {
		return "Hello" + name, true
	}
}

func fibonachi(i int) int {
	switch i {
	case 0:
		return 0
	case 1:
		return 1
	}
	return fibonachi(i-1) + fibonachi(i-2)

}

func sum(a, b int) int {
	return a + b
}

func withPoin(i *int) {
	*i += 1
}
