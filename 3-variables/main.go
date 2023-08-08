package main

import (
	"fmt"
	"reflect"
)

const PI float32 = 3.14

const SQUARE_OF_PI1 = PI * PI //This is a valid const
const SQUARE_OF_PI2 = 3.14 * 3.14

var pi float32 = 3.14

// const SQUARE_OF_PI3=pi*pi  //this is not valid as pi is a variable and can change value at runtime

const (
	DAYS    = 7
	MONTHS  = 12
	HOURS   = 24
	MINS    = 60
	SECONDS = 60
)

func main() {
	// PI=3.143
	fmt.Println(PI, SQUARE_OF_PI1, SQUARE_OF_PI2, reflect.TypeOf(SQUARE_OF_PI1))
}
