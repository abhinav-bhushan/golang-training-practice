package main

import "fmt"

func task8() {
	age := 10
	height := 120
	state := "UP"
	gender := "F"

	switch {
	case state == "Karnataka" && (gender == "F" || gender == "M" && height < 110 && age < 5):
		fmt.Println("No ticket")

	case state == "AP" && (gender == "F" && height < 110 && age < 5 || gender == "M" && height < 110 && age < 5):
		fmt.Println("No ticket")

	case state == "Delhi" && (gender == "F" || gender == "M" && height < 130 && age < 7):
		fmt.Println("No ticket")

	case state == "UP" && (gender == "F" && height < 120 && age < 6 || gender == "M" && height < 120 && age < 6):
		fmt.Println("No ticket")

	default:
		fmt.Println("Full ticket")

	}

}
