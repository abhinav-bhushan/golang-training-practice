package main

import "fmt"

func main() {
	defer fmt.Println("this is the end of main")

	defer func() {
		defer fmt.Println("this is the end of func")
		fmt.Println("Hello Victoria Secrets & Co US team")

	}()

	defer fmt.Println("Where will i be called")
	fmt.Println("Hello Victoria Secrets & Co")
	fmt.Println("Hello Victoria Secrets & Co Indian team")

}
