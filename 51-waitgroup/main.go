package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	//wg:&sync.WaitGroup{}
	i := 1
	wg.Add(103)
	//wg.Add(1)

	go func() {
		fmt.Println("Hello World")
		wg.Done()
	}()

	go func() {
		defer wg.Done()
		for i = 1; i <= 101; i++ {
			go func(num int) {
				//wg.Add(1)
				defer wg.Done()
				if i%2 == 0 {
					fmt.Println(num, "is even")
				} else {
					fmt.Println(num, "is odd")
				}
			}(i)
		}
	}()

	wg.Wait()
}
