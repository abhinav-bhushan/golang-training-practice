package main

import "fmt"

func main() {
	ch := make(chan int)
	notify := make(chan struct{})

	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i * i
		}
	}()

	go func() {
		for i := 1; i <= 100; i++ {
			fmt.Println("Reciever--", <-ch)
		}
		notify <- struct{}{} //instance of a struct,no data but empty instance
	}()

	st := <-notify
	fmt.Println("Value received from empty struct", st)

}
