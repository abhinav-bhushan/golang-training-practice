package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice1 := make([]int, 100) //created a slie

	for i, _ := range slice1 {
		slice1[i] = rand.Intn(50) //assigning values to the slice using rand function
	}

	fmt.Println(slice1)           //printing values of the slice
	mymap := make(map[int]uint16) //create a map as int and uint16

	for _, v := range slice1 { //iterating the slice1
		//in the map key is slice's value and map is count of that value in the slice
		val, ok := mymap[v] //map returns val and ok ...if ok is true that means already keye xists
		if ok {
			mymap[v] = val + 1 //if key exists value is incremented
		} else {
			mymap[v] = 1 //if key does not that means first time,then gives 1 as a default value
		}

	}

	fmt.Println(mymap)

	//func DeleteMap   key

}
