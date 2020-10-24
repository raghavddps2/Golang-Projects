package main

import "fmt"

type contactInfo struct {
	email string
	zip int
}

type person2 struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	contactRaghav := contactInfo{email: "raghav.ddps2@gmail.com", zip: 251001}
	raghav := person2 {firstName: "Raghav", lastName: "Maheshwari", contact : contactRaghav}
	shraddha := person2 {
		firstName: "Raghav",
		lastName: "Maheshwari",
		contact: contactInfo{
			email: "raghav.ddps2@gmail.com",
			zip: 251001,
		},
	}
	fmt.Println(raghav.contact.zip)
	fmt.Println(raghav.contact.email)

	fmt.Println(shraddha.lastName)
	fmt.Println(shraddha.contact.email)
	//This is simply pass by value
	shraddha.printPersonInfo()
	fmt.Println(shraddha.contact.zip)

	//As it doesn't change, it clearly tells it is pass by value and not pass  y reference
	shraddha.updateName("Bholu")
	fmt.Println(shraddha.firstName)
	// Even passing it in a function is simply pass by value and not pass by reference
	updateName(shraddha,"Bholu")
	fmt.Println(shraddha.firstName)
}

/**
	receiver as structs
 */
func (p person2) printPersonInfo() {
	fmt.Println(p.contact)
	fmt.Println(p.lastName)
	fmt.Println(p.firstName)
	p.contact.zip = 560054
}

func (p person2) updateName(newFirstName string){
	p.firstName = newFirstName
}

func updateName(p person2,name string){
	p.firstName = name
}

