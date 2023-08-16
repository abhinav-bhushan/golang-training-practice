package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go fmt.Println("Hello Victoria Secrets & Co")

	/*go func() {
		for i := 1; i <= 1000; i++ {
			num := rand.Intn(10000)
			go isEven(num)
		}
	}() */

	go func() {
		for i := 1; i <= 1000; i++ {
			go isEven(i)
		}
	}()
	go abh()

	time.Sleep(time.Second * 2) //this is not a correct way

	// runtime.Goexit() //terminates the goroutine of caller.if called in main,it waits for all other gorutines to complete its execution but the downside is that ultimately the programme crashes
	//When to use,whe u  come to the last line,you know it's gonna fail.will be used in http later
}

func isEven(num int) {
	if num%2 == 0 {
		fmt.Println(num, "is Even")
	} else {
		fmt.Println(num, "is Odd")
	}
}
func abh() {
	fmt.Println("It got executed")
	runtime.Goexit()

}

func Fatal(msg string, code int) {
	fmt.Println()
	//os.exit()
}
