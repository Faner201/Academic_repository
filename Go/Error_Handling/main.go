package main

import (
	"fmt"
	"strconv"
)

func main() {
	// err := returnError()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	result, err := returnError(true)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

type specialError struct {
	errorMessage string
	errorCode    int
}

func (s specialError) Error() string {
	return s.errorMessage + " " + strconv.Itoa(s.errorCode)
}

// func returnError() error {
// 	return errors.New("Error!")
// }
// func returnError(returnError bool) (string, error) {
// 	if returnError {
// 		return "", errors.New("Error!")
// 	}
// 	return "Random", nil
// }

func returnError(returnError bool) (string, error) {
	if returnError {
		return "", specialError{"Special Error", 123}
	}
	return "Random", nil
}
