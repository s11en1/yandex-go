package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

func (person Person) Print() {
	fmt.Printf("Name: %v \nAge: %v \nAddress: %v", person.name, person.age, person.address)
}
