package services

import (
	. "github.com/vs4vijay/matrix/pkg/interfaces"
)

type Services struct {
}

// TODO: IoC
func (services *Services) Get(serviceName string) Service {
	servicesContainer := make(map[string]Service)
	service, ok := servicesContainer[serviceName]

	if ok {
		// Do Nothing
	} else if serviceName == "HttpService" {
		// service := new(HttpService)
		// service.Init()
		servicesContainer[serviceName] = service
	} else if serviceName == "DataService" {
		service := new(DataService)
		service.Init()
		servicesContainer[serviceName] = service
	}
	return service
}
