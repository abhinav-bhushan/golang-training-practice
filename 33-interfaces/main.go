package main

func main(){

}

func Area(iarea iArea){  //interface as a parameter.It can be a dependency inject
	a:=iarea.Area()
	fmt.Println("Area :",a)
}

func Perimeter(iperimeter IPerimeter){
	p:=iperimeter.Perimeter()
	fmt.Println("Perimeter:",p)
}



type IArea interface{
	Area() float32
}

type IPerimeter interface {
	Perimeter() float32
}


