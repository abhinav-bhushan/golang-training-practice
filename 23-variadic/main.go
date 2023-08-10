package main

import "fmt"

func main() {
	s1 := sum()
	s2 := sum(10, 20)

	arr1 := [5]int{10, 11, 12, 13, 14}
	s3 := sum(arr1[:]...) //pass an array as a
	slice1 := []int{10, 11, 12, 13, 14}
	s4 := sum(slice1...)

	fmt.Println(s1, s2, s3, s4)
	fmt.Println(slice1[0:2])
	fmt.Println(slice1[3:])

	slice2 := slice1
	slice2 = append(slice1[:2], slice1[3:]...)
	fmt.Println(slice2)
	fmt.Println(slice1)

}

// the variadic parameter must be the last parameter
func sum(args ...int) int {
	_sum := 0
	for _, v := range args {
		_sum += v
	}
	return _sum

}
