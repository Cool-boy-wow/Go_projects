package storage

import "../user"

func NewStudent(name string, age int, grade int, m map[string]*user.Student) {
	m[name] = &user.Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}
}
