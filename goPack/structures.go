package main
import "fmt"
	func main(){
		rect1:=Rectangle{10,20}
		square1:=Square{10}
		fmt.Println(getArea(rect1))
		fmt.Println(getArea(square1))
	}

type Shape interface{
	area() float64
}

type Rectangle struct{
	height float64
	width float64
}

type Square struct{
	sidelength float64
}

func(r1 Rectangle) area() float64{
	return r1.height*r1.width
}

func (s1 Square) area() float64{
	return s1.sidelength*s1.sidelength
}

func getArea(shape Shape) float64{

	return shape.area()
}