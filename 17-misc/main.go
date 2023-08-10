package main

func main() {
	str1 := `"Hello World"`
	println(str1)

	a, b := 10, 20
	c := 30
	a, b, c = c, a, b
	println(a, b, c)

}
