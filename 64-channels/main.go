package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := MultipleChanOut(ch1, ch2)

	go func() {
		ch1 <- "victoria"
		ch2 <- "Secrets & Co"
	}()

	for val := range ch3 {
		fmt.Println(val)
	}
}

func MultipleChanOut(ch1, ch2 chan string) chan string {
	ch := make(chan string)
	go func() {
		ch <- <-ch1 + " " + <-ch2
		close(ch)
	}()
	return ch
}
