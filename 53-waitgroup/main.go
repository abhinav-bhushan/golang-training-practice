package main

import "fmt"

func main() {
	ch := make(chan int) //Unbuffered channel has instantiated
	go func() {
		num := <-ch
		fmt.Println(num)
	}()
	ch <- 100 //recieved value from the channel

}
