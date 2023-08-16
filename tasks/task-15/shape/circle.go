package shape

import (
	"fmt"
	"math"
)

type Circle float32

func (c Circle) Area() float32 {
	return float32(math.Pi * c * c)
}

func (c Circle) Perimeter() float32 {
	return float32(2 * math.Pi * c)
}

func (c Circle) Display() {
	fmt.Println("Circle")
	fmt.Println("-----------------")
}
