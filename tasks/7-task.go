package main

import "fmt"

func task7() {

	var state string = "Karnataka"
	var gender rune = 'M'
	var age byte = 18
	var height uint32 = 110

	if state == "Karnataka" {
		if gender == 'F' {
			fmt.Println("No ticket")
		} else {
			if height < 110 && age < 5 {
				fmt.Println("No ticket")
			} else {
				fmt.Println("Full ticket")
			}
		}
	} else if state == "Delhi" {
		if gender == 'F' {
			fmt.Println("No ticket")
		} else {
			if height < 130 && age < 7 {
				fmt.Println("No ticket")
			} else {
				fmt.Println("Full ticket")
			}
		}
	} else if state == "AP" {
		if gender == 'F' {
			if height < 110 && age < 5 {
				fmt.Println("No ticket")
			} else {
				fmt.Println("Full ticket")
			}

		} else {
			if height < 100 && age < 5 {
				fmt.Println("No ticket")
			} else {
				fmt.Println("full ticket")
			}
		}
	} else if state == "UP" {
		if gender == 'F' {
			if height < 120 && age < 6 {
				fmt.Println("No ticket")
			} else {
				fmt.Println("Full ticket")
			}

		} else {
			if height < 120 && age < 6 {
				fmt.Println("No ticket")
			} else {
				fmt.Println("full ticket")
			}
		}
	}
}
