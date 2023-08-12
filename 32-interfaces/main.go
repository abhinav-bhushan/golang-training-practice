package main

import (
	"errors"
	"fmt"
)

func main() {
	//io.writer

	fmt.Println(nil, "Hello World")

}

type Console struct{}

func (c *Console) Write(bytes []byte) (n int, err error) {
	return fmt.Println(string(bytes))
	//return 0,nil
}

type Fw struct {
	Filename string
}

func (f *Fw) Write(bytes []byte) (n int, err error) {
	if f == nil {
		return 0, errors.New("nil file object")
	}

}
