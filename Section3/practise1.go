package main

import (
	"fmt"
	"strconv"
)

type contact struct {
	zip int
	email string
	address string
}

type person struct {
	firstName string
	lastName string
	contactInfo contact
}


func main() {
	arr := []int {1,2,3,4,5,6,7,8,9,10}
	for i:= 1; i<len(arr);i++ {
		fmt.Println(arr[i])
	}

	a := 1
	b := 2
	fmt.Println(a+b)

	// Variables, if else, loops
	a = 20
	if a%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
	for i,j := range arr {
		fmt.Println(i,j)
	}
	res := multiplyBy2(arr)
	for _,j := range res {
		fmt.Println(j)
	}
	structArray()

	temp := learnAboutMaps()
	for _,j := range temp {
		fmt.Println(j)
	}
}

func structArray(){
	personsArray := []person{}

	for i:=1 ; i< 5 ; i++ {
		var cont contact
		cont.email = "abc" + strconv.Itoa(i) + "@gmail.com"
		cont.address = "abc" + strconv.Itoa(i) + "@Lane near heaven"
		cont.zip = 40000

		var pers person
		pers.contactInfo = cont
		pers.firstName = "abc" + strconv.Itoa(i)
		fmt.Println("abc"+strconv.Itoa(i))
		pers.lastName = "abc" + strconv.Itoa(i)

		personsArray = append(personsArray,pers)
	}

	for _,j := range personsArray {
		fmt.Println(j)
	}

}

func learnAboutMaps() []int {
	//Task is to find unique elements in an array
	hMap := make(map[int]int)

	arr := []int {1,1,1,2,3,3,2,2,4,5,3,3,}
	ans := []int {}
	for _,j := range arr {
		_,ok := hMap[j]
		if !ok {
			hMap[j] = 1
			ans = append(ans,j)
		}
	}
	return ans
}

func multiplyBy2(arr []int) []int {
	ans := []int {}
	for _,j := range arr {
		ans = append(ans,j*2)
	}
	return ans
}
