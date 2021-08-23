package main

import "fmt"

type ContactDetails struct {
	phoneNo string
	pincode int
}

type Person struct {
	firstName string
	lastName  string
	contact   ContactDetails
}

func main() {
	alex := Person{
		firstName: "Bhargava",
		lastName:  "Kulkarni",
		contact: ContactDetails{
			phoneNo: "8197353807",
			pincode: 586128,
		},
	}

	alexPtr := &alex
	alexPtr.print()
	alexPtr.updateFirstName("Bhavani")
	fmt.Println()
	alexPtr.print()
}

func (personPtr *Person) updateFirstName(newFirstName string) {
	(*personPtr).firstName = newFirstName
}

func (personPtr *Person) print() {
	fmt.Printf("%+v", (*personPtr))
}
