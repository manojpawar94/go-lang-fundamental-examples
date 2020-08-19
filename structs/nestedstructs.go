package main

import "fmt"

type employee struct {
	firstName string
	lastName  string
	contact
}

type contact struct {
	email  string
	mobile int
}

func example2() {
	manoj := employee{
		firstName: "Manoj",
		lastName:  "Pawar",
		contact: contact{
			mobile: 8983120926,
			email:  "mmpawar94@gmail.com",
		},
	}
	fmt.Printf("%+v", manoj)
}
