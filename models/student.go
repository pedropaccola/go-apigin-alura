package models

type Student struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func NewStudent(id int, f, l string) *Student {
	return &Student{
		ID:        id,
		FirstName: f,
		LastName:  l,
	}
}

type Students []*Student

func NewStudents() *Students {
	a := NewStudent(1, "Pedro", "Paccola")
	b := NewStudent(2, "Zeca", "Pagodinho")
	c := NewStudent(3, "Jona", "Dabe")
	return &Students{
		a, b, c,
	}
}
