package main

import "fmt"

func main() {
	//Write an anonymous function that should return number of vowels,consonants and special characters in a given string

	f2 := func(str string) (int, int, int) {
		m := make(map[string]int)
		m["vowels"] = 0
		m["consonants"] = 0
		m["special"] = 0

		for _, v := range str {
			// fmt.Println(reflect.TypeOf(v))
			switch {
			case v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u':
				m["vowels"] += 1

			case int(v) >= int('a') && int(v) <= int('z'):
				m["consonants"] += 1

			default:
				m["special"] += 1

			}
		}

		return m["vowels"], m["consonants"], m["special"]
	}
	v, c, sp := f2("hello victoria & co!")

	fmt.Printf("Special characters: %v Vowels: %v Consonants: %v\n", sp, v, c)
}
