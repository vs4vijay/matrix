package models

import (
	"matrix/pkg/services"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName string
}

func NewPerson(dataService services.DataService) *Person {
	person := &Person{
		FirstName: dataService.FirstName(),
		LastName: dataService.LastName(),
	}

	return person
}
