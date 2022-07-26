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
	p := person{
		firstName: "Victor",
		lastName:  "Batata",
		contact: contactInfo{
			email:   "victor@gmail.com",
			zipCode: 12345678,
		},
	}

	pPointer := &p
	pPointer.updateFirstName("Matheus")
	p.print()

}

func (p *person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
