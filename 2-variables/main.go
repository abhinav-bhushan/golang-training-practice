package main

import (
	"fmt"
	"reflect"
)

func main() {
	var age uint8 = 38
	fmt.Println("Age:", age, "Type of Age:", reflect.TypeOf(age))
	fmt.Printf("\nAge: %v Type of Age : %T\n", age, age)
	var num1 = 100 //Another kind of declaration
	var num2 = 25.25
	num3 := 12312 //shorthand declaration
	fmt.Println(num1, num2, num3, reflect.TypeOf(num1), reflect.TypeOf(num2), reflect.TypeOf(num3))
cd ..
	//Type inference
	var (
		num6 int     // 0
		num7 float32 // 0
		ok1  bool    // false
		str1 string  // ""
	)
	fmt.Println(num6, num7, ok1, str1)

}
