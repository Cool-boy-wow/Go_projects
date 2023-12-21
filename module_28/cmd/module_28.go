package main

import (
	"fmt"

	"../pkg/storage"
	"../pkg/user"
)

func main() {

	var name string
	age := 0
	grade := 0

	m := make(map[string]*user.Student)

	for {
		name = ""
		age = 0
		grade += 1
		fmt.Print("Введите имя студента или 'EOF' :")
		fmt.Scan(&name)
		if name != "EOF" {
			fmt.Print("Введите возраст студента:")
			fmt.Scan(&age)

			storage.NewStudent(name, age, grade, m)
		} else {
			break
		}
	}
	for k := range m {
		fmt.Println(m[k])
	}

}
