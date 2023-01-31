package main

import (
	"fmt"
)

func main() {
	type ninja struct {
		name    string
		weapons []string
		level   int
	}

	wallace := ninja{name: "Roman"}
	wallace = ninja{"Peter", []string{"Ninja Star"}, 2}
	wallace.level++
	fmt.Println(wallace.level)

	wallace.weapons = append(wallace.weapons, "Ninja sword")
	fmt.Println(wallace.weapons)

	type dojo struct {
		name  string
		ninja ninja
	}

	golangDojo := dojo{
		name:  "Golang Dojo",
		ninja: wallace,
	}

	fmt.Println(golangDojo)
	fmt.Println(golangDojo.ninja.level)
	golangDojo.ninja.level = 5
	fmt.Println(golangDojo.ninja.level)
	fmt.Println(wallace.level)

	type addressedDojo struct {
		name  string
		ninja *ninja
	}

	addreserd := addressedDojo{
		"Addressed Golang Dojo",
		&wallace,
	}

	fmt.Println(addreserd)
	fmt.Println(*addreserd.ninja)
	addreserd.ninja.level = 4
	fmt.Println(addreserd.ninja.level)
	fmt.Println(wallace.level)

	jonyPointer := new(ninja)
	fmt.Println(jonyPointer)
	fmt.Println(*jonyPointer)

	jonyPointer.name = wallace.name
	jonyPointer.weapons = []string{"Ninja Star"}
	jonyPointer.level = 3
	fmt.Println(*jonyPointer)

	intern := ninjaIternal{
		"Iternal",
		5,
	}
	intern.ninjaHello()
}

type ninjaIternal struct {
	name  string
	level int
}

func (ninjaIternal) ninjaHello() {
	fmt.Println("Hello Ninja")
}
