package main

import "fmt"

/*Student struct*/
type Student struct {
	firstName string
	lastName  string
	contact   Contact
}

/*Contact struct*/
type Contact struct {
	email  string
	mobile int
}

func example2() {
	manoj := Student{
		firstName: "Manoj",
		lastName:  "Pawar",
		contact: Contact{
			mobile: 8983120926,
			email:  "mmpawar94@gmail.com",
		},
	}
	fmt.Printf("%+v", manoj)
}
