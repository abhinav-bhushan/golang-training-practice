package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	data, ok := os.ReadFile("file.txt")

	if ok == nil {
		errors.New("File not found")
	}
	fmt.Println(string(data))

}
