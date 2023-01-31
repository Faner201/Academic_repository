package main

import (
	"fmt"
)

func main() {
	types := []string{
		"Home",
		"Job",
		"School",
		"Bar",
		"Gym",
	}

	fmt.Println(types)

	for _, typ := range types {
		switch typ {
		case "Home":
			greeting("I'm hungry")
		case "Job", "School":
			greeting("Wheater is greate today")
		case "Bar":
			greeting("Hey")
		}
	}

}
func greeting(greeting string) {
	fmt.Println(greeting)
}
