package main

import (
	"fmt"
	"math"
	"math/rand"
)

func task9() {
	var arr [10]int

	for i, _ := range arr {
		arr[i] = rand.Intn(30)

	}

	fmt.Println(arr)
	min, max := getBigAndSmall(arr)

	fmt.Printf("The minimum valiue in array is :%v and maximum is %v", min, max)

}

func getBigAndSmall(arr [10]int) (int, int) {
	min := math.MaxInt
	max := math.MinInt

	for _, v := range arr {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	return min, max

}
