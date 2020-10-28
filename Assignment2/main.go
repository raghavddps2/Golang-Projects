package main

import (
	"fmt"
)

//Shape interface
type shape interface{
	getArea() float64
}

//triangle type that implements the shape interface by defining the printArea function
type triangle struct{
	height float64
	base float64
}

//square type that implements the shape interface by defining the printArea function
type square struct{
	sideLength float64
}

//Defining the printArea() function for thr triangle type
func (tri triangle) getArea() float64 {
	ans := (tri.base*tri.height)/2
	//fmt.Println(ans)
	return ans
}

//Defining the printArea() for the square type
func (squ square) getArea() float64 {
	ans := (squ.sideLength*squ.sideLength)
	//fmt.Println(ans)
	return ans
}

func printArea(s shape) {
	//Here it will determine the struct from the interface
	fmt.Println(s.getArea())
}

//main function for testing
func main(){
	tri := triangle{
		height: 12.0,
		base: 5.0,
	}
	printArea(tri)

	squ := square{
		sideLength: 5.0,
	}
	printArea(squ)
}

