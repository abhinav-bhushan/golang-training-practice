package main

import "fmt"

func main() {
	greet()
	greetBy("Abhinav")

	a1 := areaOfRect(10.23, 13.45)
	fmt.Println(a1)
	a2, p2 := Rect(10.23, 13.45)
	fmt.Println(a2, p2)
	_, p2 = Rect(10.23, 13.45)
	a2, _ = Rect(10.23, 13.45)

	num1 := 100
	incr(num1)

	fmt.Println(num1)

}

func greet() {
	fmt.Println("hello Victoria Secrets")

}

func greetBy(name string) {
	fmt.Println("Hello", name)
}

func greetByM(msg, name string) {
	fmt.Println(msg, name)
}

func areaOfRect(l, b float32) float32 {
	return l * b
}

func Rect(l, b float32) (area float32, perimeter float32) {
	area = l * b
	perimeter = 2 * (l + b)
	return area, perimeter
}

func incr(num int) {
	num++
}
