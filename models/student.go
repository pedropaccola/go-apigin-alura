package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func NewStudent(f, l string) *Student {
	return &Student{
		FirstName: f,
		LastName:  l,
	}
}

type Students []*Student

func NewStudents() *Students {
	a := NewStudent("Pedro", "Paccola")
	b := NewStudent("Zeca", "Pagodinho")
	c := NewStudent("Jona", "Dabe")
	return &Students{
		a, b, c,
	}
}
