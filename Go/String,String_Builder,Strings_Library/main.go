package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello World"

	fmt.Println(len(s))
	fmt.Println(s[0])
	fmt.Printf("%c", s[0])
	fmt.Println(s[0:6])

	s += " Again"

	abc := "abc"
	b := []byte(abc)
	fmt.Printf("%s %s", abc, b)
	abc2 := string(b)
	fmt.Println(abc2)

	// String Library

	fmt.Println(strings.ToUpper("Hello World"))
	fmt.Println(strings.ToLower("Hello World"))
	fmt.Println(strings.HasPrefix("Hello World", "Hello"))
	fmt.Println(strings.HasSuffix("Hello World", "Hello"))
	fmt.Println(strings.Replace("Hello World", "word", "Hi", 1))

	//String Builder

	var sb strings.Builder
	sb.WriteString("Hello")
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	sb.Grow(100)
	fmt.Println(sb.Cap())

	sb.Write([]byte{62, 62, 62})
	fmt.Println(sb.String())
}
