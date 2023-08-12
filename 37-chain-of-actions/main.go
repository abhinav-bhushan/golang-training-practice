package main

import "fmt"

type ICalculator interface {
	Add(any) ICalculator
	Sub(any) ICalculator
	Mul(any) ICalculator
	Get() any
}

type AnyType struct {
	Result float64
}

func (i *AnyType) Add(num any) ICalculator {
	switch num.(type) {
	case int:
		i.Result = i.Result + float64(num.(int))
	case float32:
		i.Result = i.Result + float64(num.(float32))
	case float64:
		i.Result = i.Result + num.(float64)
	case int32:
		i.Result = i.Result + float64(num.(int32))
	}
	return i
}

func (i *AnyType) Sub(num any) ICalculator {
	switch num.(type) {
	case int:
		i.Result = i.Result - float64(num.(int))
	case float32:
		i.Result = i.Result - float64(num.(float32))
	case float64:
		i.Result = i.Result - num.(float64)
	case int32:
		i.Result = i.Result - float64(num.(int32))
	}
	return i
}

func (i *AnyType) Mul(num any) ICalculator {
	switch num.(type) {
	case int:
		i.Result = i.Result * float64(num.(int))
	case float32:
		i.Result = i.Result * float64(num.(float32))
	case float64:
		i.Result = i.Result * num.(float64)
	case int32:
		i.Result = i.Result * float64(num.(int32))
	}
	return i
}

func (i *AnyType) Get() any {
	var result any = i.Result
	return result
}

func main() {
	var icalc ICalculator

	icalc = new(AnyType)

	r1 := icalc.Add(10).Add(20).Add(20.2).Add(30).Add(30).Get()
	fmt.Println("Result r1:", r1)

	icalc = new(AnyType)
	r2 := icalc.Add(10).Sub(20).Add(20).Mul(2).Get()
	fmt.Println("Result r2:", r2)

}
