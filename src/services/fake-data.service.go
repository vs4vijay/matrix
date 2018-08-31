package services

import (
	"github.com/icrowley/fake"
)

type FakeDataService struct {
}

func(fakeDataService *FakeDataService) Init() {
}



// Fake Data Generate Functions
func(c FakeDataService) FirstName() string {
	return fake.FirstName()
}

func(c FakeDataService) LastName() string {
	return fake.LastName()
}

func(c FakeDataService) Phone() string {
	return fake.Phone()
}

func(c FakeDataService) EmailAddress() string {
	return fake.EmailAddress()
}

func(c FakeDataService) Company() string {
	return fake.Company()
}

func(c FakeDataService) Country() string {
	return fake.Country()
}

func(c FakeDataService) State() string {
	return fake.State()
}

func(c FakeDataService) Zip() string {
	return fake.Zip()
}
