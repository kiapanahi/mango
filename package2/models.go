package package2

import "fmt"

// Model1 - sample model
type Model1 struct {
	Name   string
	Family string
	Age    int
}

// NewModel1 create an instance of Model1
func NewModel1(name string, family string, age int) Model1 {
	return Model1{
		Name:   name,
		Family: family,
		Age:    age,
	}
}
func (m Model1) String() string {
	return fmt.Sprintf("{name: %s, family: %s, age: %v}", m.Name, m.Family, m.Age)
}
