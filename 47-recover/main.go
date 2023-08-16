package main

import "fmt"

func main() {
	defer Recoverme() //during the panic defer will recover but the rest of the call from 12-15 won't execute
	fn1 := "Victoria"

	str1 := GetFullname(&fn1, nil) //panics and returns nil but due to recover the stack is recovered and this function is only paniced not the remianing main function call
	// fmt.Println(*str1)             //since str1 is nil as paniced function returns nil so defrefrencing nil will give error
	func() {
		defer Recoverme()
		fmt.Println(*str1)

	}()
	ln1 := "Secrets & Co"
	str2 := GetFullname(&fn1, &ln1)

	fmt.Println(*str2)

}
func GetFullname(firstName, lastName *string) *string {
	//defer fmt.Println("Is there any panic")
	defer Recoverme() //when there is a panic it jumps to the recoverme function

	if firstName == nil || *firstName == "" {
		panic("firstname cannot be nil")
	}
	if lastName == nil || *lastName == "" {
		panic("lastname cannot be nil") //checks for defer statements
	}

	str := *firstName + *lastName
	return &str
}

func Recoverme() {
	if pan := recover(); pan != nil { //this works only when there is a panic
		fmt.Println(pan, "I will recover the whole stack except the paniced function")
	}

}
