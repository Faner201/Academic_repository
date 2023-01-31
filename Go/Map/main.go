package main

import (
	"fmt"
)

func main() {
	var m map[string]string
	m = map[string]string{}
	fmt.Println(len(m))
	m = make(map[string]string, 5)

	m = map[string]string{"Wallace": "Programmer"}
	fmt.Println(m)

	m["Johny"] = "Not a programmer"
	fmt.Println(m)
	m["Johny"] = "Programmer to be"
	fmt.Println(m["Johny"])

	delete(m, "Johny")
	fmt.Println(m)

	m["Wallace"] += " Ninja"
	m["Johny"] = "Programmer"

	for name, title := range m {
		fmt.Println(name, title)
	}

	title, ok := m["Johny"]
	if ok {
		fmt.Println(title)
	} else {
		fmt.Println("Didn't find Johny")
	}
}
