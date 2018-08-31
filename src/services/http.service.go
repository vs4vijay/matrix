package services

import (
	"log"
	"fmt"
	// "io"
	"net/http"
	// "net/url"

	)

type HttpService struct {
	ServerPort string
}

func(httpService *HttpService) Init() {
	httpService.ServerPort = "9999"
}

func(httpService HttpService) MakeServer(route string) {
	hostAndPort := ":" + httpService.ServerPort
	log.Println("Starting Webserver at", hostAndPort)

	http.HandleFunc(route, func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rw, "Go Web Server by Viz on %s", route)
		log.Println("Request", req)

		// rw <- comm
	})
	http.ListenAndServe(hostAndPort, nil)
}

// func(httpService HttpService) Request(url string, method string, payload *io.Reader) {
// 	req, err := http.NewRequest(method, url, nil)
// }

// func(httpService HttpService) Post(url string, contentType string, payload io.Reader) {
// 	resp, err := http.Post(url, contentType, payload)
// }

// func(httpService HttpService) PostForm(url string, formData url.Values) {
// 	resp, err := http.PostForm(url, formData)
// }



