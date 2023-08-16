package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	wg.Add(1) //main is also included
	wg.Add(1)

	go func() {
		fmt.Println("Hello Victoria Secrets & Co USA")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		fmt.Println("Hello Victoria Secrets & Co India")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		fmt.Println("Hello Victoria Secrets & Co Srilanka")
		wg.Done()
	}()

}
