package main

import (
	"fmt"
)

func main() {
	//Shorthand declaration
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [...]int{10, 11, 12, 13, 14}

	var arr3 [5]int = arr1
	fmt.Println(arr1, arr2, arr3)

	// var arr4 [4]int=arr2  //the length of the arrays is different
	var arr4 [4]int
	//this works if arr4 is smaller than arr2

	for i := range arr4 {
		if len(arr2) <= len(arr4) {
			arr4[i] = arr2[i]
		}
	}

	//Multi dimensional arrays

	arr2d1 := [2][2]int{{1, 2}, {3, 4}}

	for _, arrd1 := range arr2d1 {
		for _, v := range arrd1 {
			fmt.Print(v, " ")
		}
	}

	arr3d1 := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
	for _, arrd1 := range arr3d1 {
		for _, arrd2 := range arrd1 {
			for _, v := range arrd2 {
				fmt.Print(v, " ")
			}
		}
	}

	arr4d1 := [2][2][2][3]int{{{{1, 2, 12}, {3, 4, 34}}, {{5, 6, 56}, {7, 8, 78}}}, {{{1, 2, 12}, {3, 4, 34}}, {{5, 6, 56}, {7, 8, 78}}}}
	fmt.Println(arr4d1)

}
