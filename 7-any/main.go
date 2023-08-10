package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var iface1 interface //Object obj
	var iface1 any      //empty interface-u can give int,float,bool,string
	fmt.Println(iface1) //type interface gives nil
	fmt.Println("Value of ", iface1, "Type of", reflect.TypeOf(iface1))

	iface1 = 1000 //can directly assign a value
	var num1 int = iface1.(int)
	fmt.Println(num1)
	// fmt.Println("Value of ", iface1, "Type of", reflect.TypeOf(iface1))

	var float1 float32 = 10.1321
	iface1 = float1
	float2 := iface1.(float32)
	fmt.Println(float2)
	fmt.Println("Value of ", iface1, "Type of", reflect.TypeOf(iface1))

	var ok1 bool = true
	iface1 = ok1
	ok2 := iface1.(bool)
	fmt.Println(ok2)
	fmt.Println("Value of ", iface1, "Type of", reflect.TypeOf(iface1))

	var str1 string = "Victoria Secret & Co"
	iface1 = str1
	str2 := iface1.(string)
	fmt.Println(str2)
	fmt.Println("Value of ", iface1, "Type of", reflect.TypeOf(iface1))
	var ok3 bool = iface1.(bool)
	fmt.Println(ok3)

}
