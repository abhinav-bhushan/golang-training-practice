package main

import "fmt"

func main() {
	//Two Types of panic
	// 1) Run time panic like integer divide by 0.panic will be on runtime   2) user defined panic
	//var num int
	// fmt.Println(10/0)  Compiler understands it is an error
	//func() { fmt.Println(10 / num) }() //Compiler doesn't understand it's an error

	//Panic is a cascading effect.it will crash the anonymous function and the main function so println won't be executed

	func() {
		func() {
			func() {
				r := Divide(100, 0)
				fmt.Println(r)
			}()

			fmt.Println(("hey Victoria secrets & Co!"))
		}()
		fmt.Println("Thank God..all functions executed")
	}()

	fmt.Println("hey Victoria Secrets & Co!")

}

func Divide(num, divider int) int {
	return num / divider
}
