package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	d, err := os.ReadFile("file.txt")
	if d != nil {
		fmt.Errorf("%v", err)
	}
	data := string(d)
	ch := make(chan string)

	//Go Routine to send lines of the file
	go func() {
		stringsSlice := strings.Split(data, "\n")
		for _, line := range stringsSlice {
			if line != "" {
				ch <- line
			}
		}

		close(ch) //necessary to close the channel if there's a reciever waiting like for range(ch).
	}()

	//Process the data for consonants,vowels and special character
	/*numChar,numVowels,numConsonants,numSpecial:=processChars(ch,data)

	fmt.Println("The number of characters are: ",numChar)
	fmt.Println("The number of vowels are: ",numVowels)
	fmt.Println("The number of consonants are: ",numConsonants)
	fmt.Println("The number of special characters are: ",numSpecial) */
	for line := range ch {
		cha, vo, c, s := processChars(line)
		fmt.Println("--------------")
		fmt.Println("The line is: ", line)
		fmt.Println("The number of characters are: ", cha)
		fmt.Println("The number of vowels are: ", vo)
		fmt.Println("The number of consonants are: ", c)
		fmt.Println("The number of special characters are: ", s)
		fmt.Println("--------------")

	}
}

func processChars(line string) (cha, vo, c, s int) {
	//passing for number of characters
	cha = len(strings.Split(line," "))
	s = 0
	for _, v := range strings.ToLower(line) {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			vo += 1
		} else if int(v) >= int('a') && int(v) <= int('z') {
			c += 1
		} else if int(v) < 48 || int(v) > 57 {
			s += 1
		}
	}
	return

}
