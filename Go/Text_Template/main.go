package main

import (
	"fmt"
	"os"
	"text/template"
)

type secret struct {
	Name   string
	Secret string
}

func main() {

	secret := secret{
		"Ilua",
		"I love GO",
	}

	temmplatePath := "/Users/fanfurick/Documents/GitHub/Academic_repository/Go/Text_Template/template.txt"

	t, err := template.New("template.txt").ParseFiles(temmplatePath)
	if err != nil {
		fmt.Print(err)
	}

	err = t.Execute(os.Stdout, secret)
	if err != nil {
		fmt.Print(err)
	}

}
