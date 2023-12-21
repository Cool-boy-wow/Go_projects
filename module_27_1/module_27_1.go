package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Grade int
}

func NewStudent(name string, age int, grade int, m map[string]*Student) {
	m[name] = &Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}
}

func main() {

	var name string
	age := 0
	grade := 0

	m := make(map[string]*Student)

	for {
		name = ""
		age = 0
		grade += 1
		fmt.Print("Введите имя студента:")
		fmt.Scan(&name)
		if name != "EOF" {
			fmt.Print("Введите возраст студента:")
			fmt.Scan(&age)

			NewStudent(name, age, grade, m)
		} else {
			break
		}
	}
	for k := range m {
		fmt.Println(m[k])
	}

}
