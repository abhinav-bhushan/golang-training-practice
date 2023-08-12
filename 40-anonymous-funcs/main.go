package main

import (
	"fmt"
)

func main() {
	Greet()

	func() {
		str := ("Hello Victoria Secrets & Co! USA")
		fmt.Println(str)
	}() //Executor

	func(str string) {
		fmt.Println(str)

	}("Hello Victoria Secrets & Co Europe!")

	rstr := func(str string) string {
		reverse := ""

		for _, v := range str {
			// fmt.Println(reflect.TypeOf(v))
			reverse = string(v) + reverse
		}
		return reverse
	}("Hello Victoria Secrets & Co")
	fmt.Println("Reverse:", rstr)

	f1 := func(str string) string {
		reverse := ""

		for _, v := range str {
			// fmt.Println(reflect.TypeOf(v))
			reverse = string(v) + reverse
		}
		return reverse

	}

	rstr = f1("Hello Victoria Secrets & Co")
	fmt.Println("Reverse:", rstr)

	//Write an anonymous function that should return number of vowels,consonants and special characters in a given string

	f2 := func(str string) (int, int, int) {
		m := make(map[string]int)
		m["vowels"] = 0
		m["consonants"] = 0
		m["special"] = 0

		for _, v := range str {
			// fmt.Println(reflect.TypeOf(v))
			switch {
			case v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u':
				m["vowels"] += 1

			case int(v) >= int('a') && int(v) <= int('z'):
				m["consonants"] += 1

			default:
				m["special"] += 1

			}
		}

		return m["vowels"], m["consonants"], m["special"]
	}
	v, c, sp := f2("hello victoria & co!")

	fmt.Printf("Special characters: %v Vowels: %v Consonants: %v\n", sp, v, c)
	sum := Calc(10, 20, func(a1, b1 int) int { //Passing anonymous function as Closure.Calc is the caller function
		return a1 + b1
	})
	fmt.Println("Sum:", sum)
	sub := Calc(10, 20, Sub) //Passing named function as closure .Calc is the caller function
	fmt.Println("Sub:", sub)

	fnm := func(a, b int) int { //fnm is function variable
		return a * b
	}

	mul := Calc(10, 20, fnm) //Passing function variable as Closure.Calc is the caller function
	fmt.Println("Mul:", mul)

	//instatiate the map
	fnmap = make(map[string]func(int, int) int)
	//give keys and values

	operations := []string{"add", "sub", "mul", "div"}
	a := func(a, b int) int {
		return a + b
	}

	s := func(a, b int) int {
		return a - b
	}

	m := func(a, b int) int {
		return a * b
	}

	d := func(a, b int) int {
		return a / b
	}

	for _, v := range operations {
		switch {
		case v == "add":
			fnmap[v] = a
		case v == "sub":
			fnmap[v] = s
		case v == "mul":
			fnmap[v] = m
		case v == "div":
			fnmap[v] = d
		}
	}

	for k, fn := range fnmap {
		r := fn(200, 100)
		fmt.Println(k, ":", r)
	}

	/*fmt.Println("sum:", fnmap["add"](10, 20))
	fmt.Println("sub:", fnmap["sub"](10, 20))
	fmt.Println("mul:", fnmap["mul"](10, 30))
	fmt.Println("div:", fnmap["div"](20, 10))*/

}

//What is an anonymous function

func Greet() {
	fmt.Println("Hello Victoria Secret & Co! Bangalore")
}

func Calc(a, b int, fn func(a1, b1 int) int /*this is a closure.you can use  a and b inside fn as their scope is whole function*/) int {
	return fn(a, b)
}

func Sub(a, b int) int {
	c := a - b
	return c
}

var fnmap map[string]func(int, int) int
