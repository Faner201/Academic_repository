package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, _ := os.Create("file.txt")
	writer := io.Writer(file)
	n, err := writer.Write([]byte("Hello Ninja"))
	fmt.Println(n, err)
	n, err = io.WriteString(writer, "!")
	log.Println(n, err)
	file.Close()

	file, _ = os.Open("file.txt")
	reader := io.Reader(file)
	// buffer := make([]byte, 10)
	// reader.Read(buffer)
	// file.Close()
	// log.Printf("Reader n = {%v}, err = {%v}, buffer = {%v}\n", n, err, string(buffer))
	buffer := make([]byte, 2)
	for {
		n, err := reader.Read(buffer)
		log.Printf("Reader n = {%v}, err = {%v}, buffer = {%v}\n", n, err, string(buffer))
		if err != nil {
			break
		}
	}
	file.Close()

	file, _ = os.Open("file.txt")
	reader = io.Reader(file)
	buffer, err = io.ReadAll(reader)
	log.Printf("ReadAll buffer = {%v}, err = {%v}\n", string(buffer), err)

	file, _ = os.Open("file.txt")
	reader = io.Reader(file)
	buffer, err = io.ReadAll(reader)
	log.Printf("ReadAll buffer = {%v}, err = {%v}\n", string(buffer), err)
	seeker := reader.(io.Seeker)
	seeker.Seek(0, io.SeekStart)
	buffer, err = io.ReadAll(reader)
	log.Printf("ReadAll buffer = {%v}, err = {%v}\n", string(buffer), err)
}
