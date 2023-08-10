package main

import "fmt"

func main() {

	//Check whether if value is divisible by 2,3 and 6
	num := 12

	if num%2 == 0 {
		fmt.Println("Number is divisble by 2")

	} else {
		fmt.Println(num, "is not divisible by 2")
	}

	if num = 12; num%2 == 0 && num%3 == 0 && num%6 == 0 { //can declare variables in if statements
		fmt.Println(num, "is divisible by 2,3 and 6")
	} else {
		fmt.Println(num, "is divisible by 2,3 and 6")

	}

	value := 55

	//If value  <50 --Grade C

	// If value>50 and <100-Grade B

	// If value>100 -Grade A

	if value >= 0 && value < 50 {
		println("Grade C")
	} else if value > 50 && value < 100 {
		println("Grade B")
	} else {
		println("Grade A")
	}

}
