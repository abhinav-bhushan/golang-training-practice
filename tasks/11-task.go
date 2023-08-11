package main

import (
	"fmt"
	"math/rand"
)

func task11() {
	var arr2d [3][3]int

	for i, arrd1 := range arr2d {
		for j, _ := range arrd1 {
			arr2d[i][j] = rand.Intn(100)
		}
	}

	fmt.Println("Before transpose")
	fmt.Println(arr2d)
	fmt.Println("After transpose")
	transpose(&arr2d)
	fmt.Println(arr2d)

}

func transpose(arr *[3][3]int) {
	var check [3][3]bool
	for i, arr1d := range *arr {
		for j, _ := range arr1d {
			if check[i][j] == false {
				(*arr)[i][j], (*arr)[j][i] = (*arr)[j][i], (*arr)[i][j]
				check[i][j] = true
				check[j][i] = true

			}

		}
	}
}
