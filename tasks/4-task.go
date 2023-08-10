package main

import "fmt"

// Global variables
var num int = 10

func task4() {
	//Accesing global variables
	fmt.Println("Global variable num value:", num)
	{
		//local variable only present within this scope
		var num int = 30
		fmt.Println("Local variable num value:", num)

	}

}
