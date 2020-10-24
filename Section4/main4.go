package main

import "fmt"

/**

	How RAM works
		we create a struct of type person and we assign it to the variable name jim
		Go find and places the struct at an address of 0001
		the variable jim points simply to that struct

	When we call jim.updateName(), what basically happens is the new struct is passedByValue and not passedByReference
	A copy is created and that is passed, so the actual values don't change, instead a copy is changed.

	But we defintely need to change things by passed by reference. So, how are we going to do this ?
	So, if we want to pass
 */

type contact1 struct {
	email string
	zipCode int
}
type person3 struct {
	fname string
	lname string
	contactInfo contact1
}
func main(){

	raghav := person3{
		"Raghav",
		"Maheshwari",
		contact1{
			email: "raghav.ddps2@gmail.com",
			zipCode: 251001,
		},
	}

	//Now, we need to update zip code of raghav
	/**
		This line of code below helps us to assign the address of the jim to the new variable name.
		This asks to give the memory address of the struct object we created

		We don't necessarily have to do this. If we define a receiver as a pointer, we don't necesarity need to call it
		with a pointer on;y, this is a shortcut in GO!
	 */
	raghavPointer := & raghav

	fmt.Println(raghav.contactInfo.zipCode)
	raghavPointer.updateZipCode(560054)
	fmt.Println(raghav.contactInfo.zipCode)
}

/**
	Here, we dereference that using the * operator.
	And we have the reference to the actual struct
	*person3 implies we are working with a pointer to a person
 */
func (raghavPointer *person3) updateZipCode(zipCode int) {
	//Here it is an operator that changes it to the actual struct we were referring to
	(*raghavPointer).contactInfo.zipCode = zipCode
}