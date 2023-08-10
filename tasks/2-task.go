package main

import "fmt"

func task2() {
	//println - prints,adds space between operands and adds new line character
	fmt.Println("Hello", "world")

	//print-prints and stays at the same line

	fmt.Print("Hello World")

	//printf -formatted print
	var str string = "world"
	fmt.Printf("Hello %v", str)
}
