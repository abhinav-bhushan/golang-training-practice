package main

import "fmt"

func main() {
	var icalc ICalculator //nil interfaces can also be nil if nothing has been assigned
	if icalc == nil {
		fmt.Println("nil interface")
	}
	it1 := &IntType{A: 100, B: 200}
	icalc = it1
	r := icalc.Add()
	fmt.Println("Addition of IntType")
	fmt.Println(r)
	it2 := &Float32Type{A: 100., B: 300.45}
	icalc = it2
	r = icalc.Add()
	fmt.Println("Addition of Float32Type")
	fmt.Println(r)
	fmt.Println("Addition of Float64Type")
	it3 := &Float64Type{A: 1234.343, B: 24343.342}
	icalc = it3
	r = icalc.Add()
	fmt.Println(r)

}

type ICalculator interface {
	Add() any //Add 2 numbers
	//Sub() any   //Sub 2 numbers
	//Mul() any   //Mul 2 numbers
	//Div() any  //Div 2 numbers
	//Get() any  //Get result
	//Clear()    //Clear
}

type IntType struct {
	A, B   int
	Result int
}

func (i *IntType) Add() any {
	i.Result = i.A + i.B
	return i.Result
}

func (i *IntType) Sub() any {
	i.Result = i.A - i.B
	return i.Result
}

type Float32Type struct {
	A, B   float32
	Result float32
}

func (f *Float32Type) Add() any {
	f.Result = f.A + f.B
	return f.Result
}

func (f *Float32Type) Sub() any {
	f.Result = f.A - f.B
	return f.Result
}

type StringType struct {
	Str1, Str2 string
	Result     string
}

func (s *StringType) Add() any {
	s.Result = s.Str1 + s.Str2
	return s.Result
}

/*func (s *StringType) Sub()any{
	s.Result=s.Str1-s.Str2
	return s.Result
} */

type Float64Type struct {
	A, B   float64
	Result float64
}

func (f *Float64Type) Add() any {
	f.Result = f.A + f.B
	return f.Result
}

func (f *Float64Type) Sub() any {
	f.Result = f.A - f.B
	return f.Result
}
