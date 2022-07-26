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

	fmt.Printf("%+v", p)
}
