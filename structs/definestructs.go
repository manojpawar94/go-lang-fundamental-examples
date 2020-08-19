package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func example1(){
	manoj := person{"Manoj", "Pawar"}
	fmt.Println(manoj)

	raja := person{firstName: "Raja", lastName: "Sundervel"}
	fmt.Println(raja)

	var sai person
	fmt.Println(sai)
	fmt.Printf("%+v\n", sai)

	sai.firstName = "Sai"
	sai.lastName = "Amblesh"
	fmt.Printf("%+v",sai)
}