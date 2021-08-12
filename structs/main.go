package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	//leo := person{"Leo", "Carrasco"}
	//leo := person{firstName: "Leo", lastName: "Carrasco"}
	//var leo person
	//leo.firstName = "Leo"
	//leo.lastName = "Carrasco"
	//leo.contact.email = "a@a.com"
	//leo.contact.zipCode = "000"
	leo := person{
		firstName: "Leo",
		lastName:  "Carrasco",
		contact: contactInfo{
			email:   "a@a.com",
			zipCode: 000,
		},
	}

	leo.print()

	leoPointer := &leo
	leoPointer.updateFirstName("Leonardo")

	leo.print()

}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*&pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}
