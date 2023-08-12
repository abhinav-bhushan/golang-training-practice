package main

import (
	"errors"
	"fmt"
)

func task13() {

	m := make(map[uint32]uint)
	m[23443] = 3473838
	m[36473] = 346282

	fmt.Println("Invalid key delete")
	fmt.Println(m)
	_, ok1 := mydelete(m, 234)
	fmt.Println("Error:", ok1)
	fmt.Println("Valid key delete")
	mydelete(m, 23443)
	fmt.Println("Map now:", m)

	fmt.Println("Invalid Key Update")
	_, ok2 := myUpdate(m, 234, 456)
	fmt.Println("Error:", ok2)
	fmt.Println("Valid key update")
	myUpdate(m, 36473, 3467)
	fmt.Println(m)

	fmt.Println("If the map is empty")
	m3 := make(map[uint32]uint)
	_, ok3 := mydelete(m3, 23)
	_, ok4 := myUpdate(m3, 23, 2)
	fmt.Println(ok3, ok4)

}

func mydelete(m map[uint32]uint, key uint32) (bool, error) {
	//if map is nil return error
	if m == nil {
		return false, errors.New("Map is empty")
	}

	if _, ok := m[key]; ok == false {
		return false, errors.New("Map doesn't contain the key")
	} else {
		delete(m, key)
		return true, nil
	}
}

func myUpdate(m map[uint32]uint, key uint32, value uint) (bool, error) {
	//if map is nil return error
	if m == nil {
		return false, errors.New("Map is empty")
	}

	if _, ok := m[key]; ok == false {
		return false, errors.New("Map doesn't contain the key")
	} else {
		m[key] = value
		return true, nil
	}

}
