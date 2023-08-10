package main

import (
	"fmt"
	"reflect"
)

func task3() {
	const num1 = 5
	const num2 = 10

	//Addition
	const output_1 = num1 + num2
	fmt.Println(output_1, reflect.TypeOf(output_1))

	//Subtraction

	const num3 = 4565
	const num4 int = 5678
	const output_2 = num3 - num4
	fmt.Println(output_2, reflect.TypeOf(output_2))

	// multiplication
	const num5 = 3245
	const num6 = 3456
	const output_3 = num5 * num6
	fmt.Println(output_3, reflect.TypeOf(output_3))

	//Diviison

	const num7 = 1456.234
	const num8 = 23.567
	const output_4 = num7 / num8
	fmt.Println(output_4, reflect.TypeOf(output_4))

	//Modulos

	const num9 = 2345.789
	const num10 = 1234.567
	// const output_5 = math.Mod(num9, num10)

}
