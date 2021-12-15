package main

import (
	"fmt"
	"math"
)

func main() {
	//we have to implement all methods of interface
	// implement interface geometry for struct rectangle
	var rect geometry = rectangle{height: 4, width: 5}
	fmt.Println("Area of rectangle: ", rect.area())
	fmt.Println("Perimeter of rectangle: ", rect.perimeter())
	// implement interface geometry for struct rectangle
	var c geometry = circle{radius: 4}
	fmt.Println("Area of circle: ", c.area())
	fmt.Println("Perimeter of circle: ", c.perimeter())
}

type geometry interface {
	area() float64 //Declare methods / behaviours signature
	perimeter() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	height, width float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2*r.height + 2*r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
