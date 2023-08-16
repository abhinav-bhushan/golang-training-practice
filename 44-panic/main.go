package main

import "fmt"

func main() {
	fn1 := "Victoria"
	ln1 := "Secrets & Co"

	str1 := GetFullname(&fn1, &ln1)
	fmt.Println("Full Name:", str1)

	fn2, ln2 := new(string), new(string)
	str2 := GetFullname(fn2, ln2)
	fmt.Println("Full Name:", str2)

	*fn2, *ln2 = "Victoria", "Secrets & Co"
	str2 = GetFullname(fn2, ln2)
	fmt.Println("Full Name:", str2)
}

func GetFullname(firstName, lastName *string) *string {

	if firstName == nil || *firstName == "" {
		panic("firstname cannot be nil")
	}
	if lastName == nil || *lastName == "" {
		panic("lastname cannot be nil")
	}

	str := *firstName + *lastName
	return &str
}

//When to raise panic:
//Without a  sucessful execution,if your application cannot proceed further then you can consider raising a panic
//Ex: No database connection
