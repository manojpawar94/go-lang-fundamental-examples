package main

import "fmt"

/*Person struct*/
type Person struct {
	firstName string
	lastName  string
}

func example1(){
	manoj := Person{"Manoj", "Pawar"}
	fmt.Println(manoj)

	raja := Person{firstName: "Raja", lastName: "Sundervel"}
	fmt.Println(raja)

	var sai Person
	fmt.Println(sai)
	fmt.Printf("%+v\n", sai)

	sai.firstName = "Sai"
	sai.lastName = "Amblesh"
	fmt.Printf("%+v",sai)
}