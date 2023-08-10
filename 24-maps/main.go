package main

import "fmt"

func main() {
	var mymap map[string]any //declaration but not instantiation
	mymap = make(map[string]any)
	mymap["522001"] = "Guntur-1"
	mymap["522002"] = "Guntur-2"
	mymap["522086"] = "Yeshvantpur Bangalore"
	mymap["5220096"] = "Mahalakshmi Bangalore"

	for key, value := range mymap {
		fmt.Println("Key:", key, "value:", value)
	}

	_, ok := mymap["522001"] //returns value and boolean
	if ok {
		fmt.Println("Key exists")
	} else {
		fmt.Println("Key doees not ")
	}

	delete(mymap, "522001")  //delete is used to delete the key in map
	delete(mymap, "5234324") //delete won't do anything if  key is not im map

	for key, value := range mymap {
		fmt.Println("Key:", key, "value:", value)
	}

}
