package main

import "fmt"

func main() {
	a := 10
	b := 20
	defer fmt.Println(add(a, b)) //this will store the variable in it's call stack ie a and b wll be the values presnet during the call of defer function and saved

	defer func() {
		c := add(a, b) //this won't save  a and b in it's call stack.a and b will be the values present at the of function which is 11 and 21 instead of 10 and 20 as they don't get saved in their call stack
		fmt.Println(a, b, c)
	}()

	a = 11
	b = 21
	c := add(a, b)
	fmt.Println(a, b, c)

}

func add(a, b int) int {
	return a + b
}
