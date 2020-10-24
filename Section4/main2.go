package main

import "fmt"

type person1 struct {
	firstName string
	lastName string
}

/**
	The zero value for different types of datatypes is
	string: ""
	int: 0
	float: 0
	bool: false
 */
func main() {
	//When we specifically want the zero values with the structures.
	var alex person1
	fmt.Println(alex)
	fmt.Println(alex.firstName,"This is an empty string")
	fmt.Println(alex.lastName,"This is an empty string")

	alex.firstName = "Raghav"
	alex.lastName = "Maheshwari"
	fmt.Println(alex)

}
