package main

func main() {

	//Check if a number is divisible by 2,4 and 8

	divisibleBy842(32)
	divisibleBy842(28)
	divisibleBy842(3)
	divisibleBy842(5)

	println("wrong results as false negative.Wrong implementation of fallthrough")
	falseNegativeImplementationOfFallthrough(28)
	falseNegativeImplementationOfFallthrough(6)

}

func divisibleBy842(num int) {
	switch {
	case num%8 == 0:
		println("Number is divisible by 8")
		fallthrough
	case num%4 == 0:
		println("Number is divisible by 4")
		fallthrough
	case num%2 == 0:
		println("Number is divisible by 2")
	default:
		println("divisible by all or none..jUST FOR DEMO PURPOSE")
	}

}

func falseNegativeImplementationOfFallthrough(num int) {
	switch {
	case num%2 == 0:
		println("Number is divisible by 2")
		fallthrough
	case num%4 == 0:
		println("Number is divisible by 4")
		fallthrough
	case num%8 == 0:
		println("Number is divisible by 8")
	}

}
