package main

import "fmt"

func main() {
	str1 := "Victoria Secrets and Co"
	str2 := "维多利亚的秘密公司"

	for i := 0; i < len(str1); i++ {
		fmt.Print(string(str1[i]), " ")
	}
	fmt.Println()
	for i := 0; i < len(str2); i++ {
		fmt.Print(string(str2[i]), " ")
	}

	for _, v := range str2 {
		fmt.Print(string(v), " ")
	}
}
