package main

import (
	"demo/fileops"
	"fmt"
)

func main() {
	/*file,err:=os.OpenFile("data.txt",1,0644)

	if err!=nil {
		fmt.Println(err)
	}*/

	fileName := new(string)
	*fileName = "dat.x"

	count, err := fileops.GetCharCount(fileName)
	if err != nil {
		fmt.Println("error:", err)
		//or do what logic is to be implemented
	} else {
		fmt.Println("Char count:", count)
	}
}
