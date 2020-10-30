package main

import "fmt"

type person struct{
	fname string
	lname string
}


type shape interface {
	getArea() int
}

type rectangle struct {
	length int
	breadth int
}

type square struct {
	side int
}

func (rect rectangle) getArea() int {
	return rect.length*rect.breadth
}
func (squ square) getArea() int {
	return squ.side*squ.side
}

func printArea(shape1 shape) {
	//Here we identify the ty[pe easily!
	//Both square and rectangle are of type struct as they have the getArea() method.
	//This is a Polymorphis (run time ploymorphism) when which types function to be called is determined at the run time.
	
	fmt.Printf("%T",shape1)
	fmt.Println(shape1.getArea())
}

type secretAgent struct {
	person person
	license bool
}


type hotdog int
func main(){
	xi := []int{2,4,6,8,10}
	for i,j := range xi {
		fmt.Println(i,j)
	}

	m := map[string]int {
		"raghav":21,
		"rahul":22,
	}

	for i,j := range m {
		fmt.Println(i,j)
	}

	var h hotdog
	h = 23
	fmt.Println(h)

	var p person
	p.fname = "raghav"
	p.lname = "maheshwari"

	fmt.Println(p)

	p1 := person{
		fname: "raghav",
		lname: "maheshwari",
	}
	fmt.Println(p1)
	sag := secretAgent{
		person: p1,
		license: false,
	}
	fmt.Println(sag)
	fmt.Println(p1.giveAge(10))
	sag.speakUp(2,"raghav")

	s1 := square{
		side: 10,
	}
	r1 := rectangle{
		length: 2,
		breadth: 2,
	}

	printArea(s1)
	printArea(r1)
}

func (secretAg secretAgent) speakUp(n int, s string) (bool,int) {
	fmt.Println(secretAg.person.lname,"is ",n)
	return true,10
}


func (p person) giveAge(n int) int {
	fname := p.fname
	lname := p.lname

	fmt.Println(fname,lname,n)
	return 2
}
