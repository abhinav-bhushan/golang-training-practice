package main

import (
	"fmt"
	"reflect"
)

func task6() {

	var (
		v1  int8
		v2  uint8
		v3  int16
		v4  uint16
		v5  int32
		v6  uint32
		v7  float32
		v8  int64
		v9  uint64
		v10 float64
		v11 int
		v12 byte
		v13 uint
		v14 rune
	)
	v1 = -34
	v2 = 125
	v3 = -3456
	v4 = 3467
	v5 = -56864
	v6 = 34672
	v7 = 2345.8930
	v8 = -2383937
	v9 = 2536382
	v10 = 25367.364739
	v11 = 234
	v12 = 255
	v13 = 47
	v14 = 'A'

	var iface1, iface2 any
	iface1 = 2
	iface2 = 3.

	sum := int(v1) + int(v2) + int(v3) + int(v4) + int(v5) + int(v6) + int(v8) + int(v9) + v11 + int(v12) + int(v13) + int(v14) + iface1.(int)
	mul := float64(v7) + float64(v9) + float64(v10) + iface2.(float64)

	fmt.Println("Addition:", sum)
	fmt.Println("Multiplication", mul)
	fmt.Println("Type of iface1:", reflect.TypeOf(iface1))
	fmt.Println("Type of iface2:", reflect.TypeOf(iface2))

}
