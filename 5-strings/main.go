package main

import "fmt"

var (
	num1 int
	num2 float32
	num3 float64
)

func main() {
	var str1 string = "Hello World"
	fmt.Println(str1)

	fmt.Println("Length of str1 ", len(str1))

	//mutating string
	str1 = "Hello Victoria Secrets& Co"
	fmt.Println("Length of str1 ", len(str1))
	num1 := 1000 //Shorthand declaration.It is int and int holds 8 bytes on 64 bit processors
	fmt.Println(num1)
	num1 = 2000 //mutating a num1
	fmt.Println(num1)
	num1 = 38388
	fmt.Println(num1)

}
