package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {

	var iface1, iface2 any

	if sum, err := add(iface1, iface2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum)
	}
}

// a and b must be of same type
func add(a, b any) (float64, error) {
	sum := 0. //If you give 0. or 0.0 sum is treated as float64

	if a == nil || b == nil {
		return 0, fmt.Errorf("a or b is nil")
	}

	if reflect.TypeOf(a) != reflect.TypeOf(b) { //Both of them are different types return an error
		return 0, errors.New("a and b must be of same type")
	}

	//Here we ensure that a and b is of same type and none of them are nil

	switch a.(type) { //this is how you need to check type of a variable which is of type any
	case int:
		sum = float64(a.(int) + b.(int))
	case uint:
		sum = float64(a.(uint) + b.(uint))
	case uint8:
		sum = float64(a.(uint8) + b.(uint8))
	case uint16:
		sum = float64(a.(uint16) + b.(uint16))
	case uint32:
		sum = float64(a.(uint32) + b.(uint32))
	case uint64:
		sum = float64(a.(uint64) + b.(uint64))
	case int8:
		sum = float64(a.(int8) + b.(int8))
	case int16:
		sum = float64(a.(int16) + b.(int16))
	case int32:
		sum = float64(a.(int32) + b.(int32))
	case int64:
		sum = float64(a.(int64) + b.(int64))

	}
	return sum, nil

}
