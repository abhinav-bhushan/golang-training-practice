package main

import "fmt"

func main() {
	var arr1 [5]int
	fmt.Println(arr1)

	var arr2 [5]bool
	fmt.Println(arr2)

	var arr3 [3]any
	fmt.Println(arr3)

	arr1[0] = 10
	arr1[1] = 11
	arr1[2] = 12
	arr1[3] = 13
	arr1[4] = 14

	fmt.Println("length of an array arr1", len(arr1))
	fmt.Println("Size of an array arr1", cap(arr1)) //There is no usage for cap of array

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i], " ")
	}

	//for range loop
	//range loop is applicable for array,slice,map,channel
	//if range loop is used on array or slice,it gives index and value

	fmt.Println("Range loop index and value")

	for index, value := range arr1 {
		fmt.Println("Index", index, "Value", value)
	}

	fmt.Println("Range loop index")

	for index, _ := range arr1 {
		fmt.Println("Index", index)
	}

	fmt.Println("Range loop value")

	for _, value := range arr1 {
		fmt.Println("Value", value)
	}

	sum := SumOfArr(arr1)
	fmt.Println("Sum of elements of an array arr1", sum)

}

func SumOfArr(arr [5]int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}
