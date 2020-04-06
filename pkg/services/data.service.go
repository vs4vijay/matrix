package services

import (
	"github.com/icrowley/fake"
)

type DataService struct {
}

func (dataService *DataService) Init() {
}

// Fake Data Generate Functions
func (c DataService) FirstName() string {
	return fake.FirstName()
}

func (c DataService) LastName() string {
	return fake.LastName()
}

func (c DataService) Phone() string {
	return fake.Phone()
}

func (c DataService) EmailAddress() string {
	return fake.EmailAddress()
}

func (c DataService) Company() string {
	return fake.Company()
}

func (c DataService) Country() string {
	return fake.Country()
}

func (c DataService) State() string {
	return fake.State()
}

func (c DataService) Zip() string {
	return fake.Zip()
}
