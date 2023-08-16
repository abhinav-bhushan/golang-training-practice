package main

import (
	"fmt"
	"reflect"
)

func main() {
	Calc("Hi", "Victoria")  //Recover will recover the stack except the paniced function
	Calc([]int{10, 20}, 10) ///Recover will recover the stack except the paniced function
	sum, sub, mul, div := Calc(10, 20)
	fmt.Println(sum, sub, mul, div)
}

func Calc(a, b any) (float64, float64, float64, float64) {
	defer RecoverMe()
	var sum, sub, mul, div float64

	switch a.(type) { //this is how you need to check type of a variable which is of type any
	case int:
		sum = float64(a.(int) + b.(int))
		mul = float64(a.(int) * b.(int))
		sub = float64(a.(int) - b.(int))
		div = float64(a.(int) / b.(int))
	case uint:
		sum = float64(a.(uint) + b.(uint))
		mul = float64(a.(uint) * b.(uint))
		sub = float64(a.(uint) - b.(uint))
		div = float64(a.(uint) + b.(uint))
	case uint8:
		sum = float64(a.(uint8) + b.(uint8))
		mul = float64(a.(uint8) * b.(uint8))
		sub = float64(a.(uint8) - b.(uint8))
		div = float64(a.(uint8) / b.(uint8))

	case uint16:
		sum = float64(a.(uint16) + b.(uint16))
		mul = float64(a.(uint16) * b.(uint16))
		sub = float64(a.(uint16) - b.(uint16))
		div = float64(a.(uint16) / b.(uint16))

	case uint32:
		sum = float64(a.(uint32) + b.(uint32))
		mul = float64(a.(uint32) * b.(uint32))
		sub = float64(a.(uint32) - b.(uint32))
		div = float64(a.(uint32) / b.(uint32))
	case uint64:
		sum = float64(a.(uint64) + b.(uint64))
		mul = float64(a.(uint64) * b.(uint64))
		sub = float64(a.(uint64) - b.(uint64))
		div = float64(a.(uint64) / b.(uint64))
	case int8:
		sum = float64(a.(int8) + b.(int8))
		mul = float64(a.(int8) * b.(int8))
		sub = float64(a.(int8) - b.(int8))
		div = float64(a.(int8) / b.(int8))

	case int16:
		sum = float64(a.(int16) + b.(int16))
		mul = float64(a.(int16) * b.(int16))
		sub = float64(a.(int16) - b.(int16))
		div = float64(a.(int16) / b.(int16))

	case int32:
		sum = float64(a.(int32) + b.(int32))
		mul = float64(a.(int32) * b.(int32))
		sub = float64(a.(int32) - b.(int32))
		div = float64(a.(int32) / b.(int32))

	case int64:
		sum = float64(a.(int64) + b.(int64))
		mul = float64(a.(int64) + b.(int64))
		sub = float64(a.(int64) - b.(int64))
		div = float64(a.(int64) / b.(int64))

	case float32:
		sum = float64(a.(float32) + b.(float32))
		mul = float64(a.(float32) * b.(float32))
		sub = float64(a.(float32) - b.(float32))
		div = float64(a.(float32) / b.(float32))

	case float64:
		sum = a.(float64) + b.(float64)
		mul = a.(float64) * b.(float64)
		sub = a.(float64) - b.(float64)
		div = a.(float64) / b.(float64)

	default:
		s := fmt.Sprintf("Invalid type detected: %v", reflect.TypeOf(a))
		panic(s)
	}

	return sum, sub, mul, div

}

func RecoverMe() {
	if pan := recover(); pan != nil {
		fmt.Println(pan, "I will recover the whole stack except the paniced function")
	}

}
