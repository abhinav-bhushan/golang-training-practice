package main

import (
	"fmt"
	"math"
)

func task12() {
	//slice
	arr := [5]int{1, 2, 3}
	fmt.Println(arr)
	//slice := arr[:] //pointing to the same arr so modifying it changes the backing array
	/* slice[0] = 2
	fmt.Println(arr) */
	secondlargest := SecondBiggest(arr[:])
	secondsmallest := SecondSmallest(arr[:])

	fmt.Printf("Second largest: %v && Second Smallest: %v", secondlargest, secondsmallest)

}

// passing array as a slice-->passig by ref equivalent
func SecondBiggest(arr []int) (secondmax int) {
	max := math.MinInt
	secondmax = max
	prevMax := max
	for i, v := range arr {
		if v > max {
			if i == 0 {
				max = v
				continue
			}
			prevMax = max
			max = v
			secondmax = int(math.Max(float64(prevMax), float64(secondmax)))

		} else if v > secondmax {
			secondmax = v

		}
	}

	return

}

func SecondSmallest(arr []int) (secondmin int) {
	min := math.MaxInt
	secondmin = min
	prevMin := min
	for _, v := range arr {
		if v < min {
			prevMin = min
			min = v
		} else if v > min && v < secondmin {
			secondmin = int(math.Min(float64(prevMin), float64(v)))
		}
	}
	return

}
