package main

import "fmt"

/**
	Data Structure - STruct - Collection of properties that are related together
	struct {
		suit : string
		value : string
	}
	At a very high level, we can think of these as structures in C
 */

/**
	Declaring a struct is like super easy. All we have to do is to define the properties inside type person struct
 */
type person struct {
	firstName string
	lastName string
}

func main() {

	//Method 1 to declare
	person1 := person {"Alex","Anderson"}
	fmt.Println(person1.firstName)
	fmt.Println(person1.lastName)
	fmt.Println(person1)
	//Method2 to declare - this helps us if we change the order of the fields, our code wont be affected
	person2 := person {firstName: "Raghav", lastName: "Maheshwari"}
	fmt.Println(person2.firstName)
	fmt.Println(person2.lastName)
	fmt.Println(person2)


}
