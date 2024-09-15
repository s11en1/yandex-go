package students

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (student Student) PrintInfo() {
	fmt.Println("Студент:", student.Name, "Возраст:", student.Age)
}
