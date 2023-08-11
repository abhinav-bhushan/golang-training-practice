package main

import (
	"fmt"
	"math/rand"
)

func task10() {
	var arr2d [3][3]int

	for i, arrd1 := range arr2d {
		for j, _ := range arrd1 {
			arr2d[i][j] = rand.Intn(100)
		}
	}

	fmt.Println(arr2d)
	for row, _ := range arr2d {
		fmt.Printf("Row %v Sum: %v\n", row+1, getSumRow(arr2d, row))

	}
	for col := 0; col < len(arr2d[0]); col++ {
		fmt.Printf("Col %v Sum: %v\n", col+1, getSumCol(arr2d, col))
	}

}

func getSumRow(arr [3][3]int, row int) (sum int) {
	for _, v := range arr[row] {
		sum += v
	}
	return
}

func getSumCol(arr [3][3]int, col int) (sum int) {
	for _, arr1d := range arr {
		sum += arr1d[col]
	}
	return
}
