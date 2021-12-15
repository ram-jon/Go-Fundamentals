//  Go methods are similar to Go function with one difference, i.e, the method contains a receiver argument in it. With the help of the receiver argument, the method can access the properties of the receiver

package main

import "fmt"

func main() {

	rect := Rectangle{
		width:  5,
		height: 3,
	}
	//calling methods with struct receiver
	area := rect.area()
	fmt.Println(area)
	p := &rect
	p.setWidth(40) //method with pointer receiver
	area = rect.area()
	fmt.Println("Area after setArea:", area)
	// calling method with non-struct
	value1 := data(23)
	value2 := data(20)
	res := value1.multiply(value2)
	fmt.Println("Final result: ", res)
}

//Methods with struct type receiver
type Rectangle struct {
	width, height int
}

func (r Rectangle) area() int {
	return r.width * r.height
}

// methods with pointer receiver
func (r *Rectangle) setWidth(w int) {
	r.width = w
}

// Methods with non struct type receiver

type data int

func (d1 data) multiply(d2 data) data {
	return d1 * d2
}
