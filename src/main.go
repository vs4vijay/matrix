package main

import (
	"fmt"
	"log"
	// . "models"
	. "services"
)

func init() {
	log.SetPrefix("LOG: ")
}

func main() {
	log.Println("Event Matrix")
	//services := new(Services)
	//httpService := services.Get("HttpService")
	httpService := new(HttpService)
	//httpService.Init()

	// comm := make(chan string)

	go httpService.MakeServer("/yo")

	// comm <- "heyy"

	fakeDataService := new(FakeDataService)
	fakeDataService.Init()

	name := fakeDataService.FirstName()

	fmt.Printf("name 1 %+v", name)

	// comm <- "hahah"
}