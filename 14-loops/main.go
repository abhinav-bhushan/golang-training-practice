package main

import (
	"errors"
	"fmt"
)

func main() {

	for i := 1; i <= 10; i++ {
		print(i, " ")

	}
	println("Even number")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			print(i, " ")
		}
	}

	println("even numbers with continue")

	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue //continue continues to the next
		}
		print(i, " ")
	}

	num:=

	//check whether a number is prime or not
	if ok, err := isPrime(0); err != nil {
		fmt.Println(err)
	} else {
		if ok {
			fmt.Println(num, "is a prime number")
		} else {
			fmt.Println(num, "is not a prime numnber")
		}
	}

	j:=1

	for {
		if j<=10{
			fmt.Println(("Victoria Secrets"))
		} else{
			break
		}
	}

}

func isPrime(num int) (bool, error) {

	if num == 0 {
		return false, errors.New("num is zero;Hence cannot decide")
	}

	if num == 1 || num == -1 {
		return true, nil
	}

	isPrime := true
	i := 2
	for ; i < num/2; i++ {

		if num%i == 0 {
			println(num, " is not a primer number\n")
			isPrime = false
			return isPrime, nil
		}
	}
	return isPrime, nil
}
