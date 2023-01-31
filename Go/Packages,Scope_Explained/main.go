package main

import (
	"fmt"
)

var five = 5

func main() {
	//Visibility in blocks

	{
		integer := 2
		fmt.Println(integer)
	}

	integer := 3

	fmt.Println(integer)

	noMain()

}

func noMain() {
	fmt.Println(five)
}
