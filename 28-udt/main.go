package main

import "fmt"

func main() {

	r1 := Rect{L: 100, B: 200}
	a1 := AreaOfRect(r1)
	fmt.Println("Area of rect r1", a1)

	a2 := r1.Area()
	fmt.Println("Area of rect r1", a2)

}

type Rect struct {
	L, B float32
	A, P float32
}

type Square struct {
	S float32
}

//to write method,need to add a reciever

func (r Rect) AreaBy() float32 {
	r.A = r.L * r.B
	return r.A
}

func (r Rect) PerimeterBy() float32 {
	r.P = r2 * (r.L + r.B)
	return r.P
}

func (r *Rect) Area() float32 {
	r.A = r.L * r.B
	return r.A
}

func (r *Rect) Perimeter() float32 {
	r.P = r2 * (r.L + r.B)
	return r.P
}

// AreaOfRect is a fubction
func AreaOfRect(r Rect) float32 {
	return r.L * r.B
}

func PerimeterOfRec(r Rect) float32 {
	return 2 * (r.L + r.B)
}
