package main

import (
	"lesson/students"
)

func main() {
	student := students.Student{Name: "Олег", Age: 2}

	student.PrintInfo()
}
