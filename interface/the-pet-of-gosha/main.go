package main

import "fmt"

type Animal interface {
	MakeSound()
}

type Dog struct{}

type Cat struct{}

func (cat Cat) MakeSound() {
	fmt.Println("Мяу!")
}

func (dog Dog) MakeSound() {
	fmt.Println("Гав!")
}
