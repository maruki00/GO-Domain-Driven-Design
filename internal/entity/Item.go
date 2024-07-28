package entity

type Person struct {
	ID   int
	Name string
	Age  int
}

func (obj Person) getId() int {
	return obj.ID
}
