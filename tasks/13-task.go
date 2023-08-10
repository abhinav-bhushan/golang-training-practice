package main

import "fmt"

func main() {

	m := make(map[uint32]uint)
	m[23443] = 3473838
	m[36473] = 346282
	temp := 0

	for k, _ := range m {
		delete(m, k)
		temp = 1
	}
	if temp == 0 {
		fmt.Println("The map doesn't contain the key")
	} else {
		fmt.Println("The key got deleted")
	}

}

/* func delete(mymap map[uint32]uint, key uint32) (bool,error) {
	if mymap==nil{


	}

}*/
