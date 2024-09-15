package main

import "fmt"

type Employee struct {
	name     string
	position string
	salary   float64
	bonus    float64
}

func (employee Employee) CalculateTotalSalary() {
	fmt.Printf("Employee: %v \nPosition: %v \nTotal Salary: %.2f", employee.name, employee.position, employee.salary+employee.bonus)
}
