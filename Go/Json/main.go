package main

import (
	"encoding/json"
	"fmt"
)

type Ninja struct {
	Name   string `json:"full_name"`
	Weapon Weapon `json:"weapon"`
	Level  int    `json:"level"`
}

type Weapon struct {
	Name  string `json:"weapon_name"`
	Level int    `json:"weapon_level"`
}

func main() {
	sForm := `{"full_name" : "Pol", "weapon" : {"weapon_name" : "Ninja Star", "weapon_level" : 5}, "level" : 2}`
	var wallace Ninja
	err := json.Unmarshal([]byte(sForm), &wallace)
	if err != nil {
		fmt.Println("Sad boy")
	}
	fmt.Println(wallace)

	sTo, err := json.Marshal(wallace)
	if err != nil {
		fmt.Println("Sad boy begin")
	}
	fmt.Printf("%s\n", sTo)
}
