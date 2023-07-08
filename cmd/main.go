package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", shortenURLHandler)
	mux.HandleFunc("/fetch/{id}", redirectURLHandler)

	err := http.ListenAndServe("localhost:8080", mux)
	log.Println("Server started on http://localhost:8080")
	if err != nil {
		panic(err)
	}

}

func shortenURLHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		request.ParseForm()

		url := request.Form.Get("url")

		response.WriteHeader(http.StatusCreated)
		response.Header().Set("content-type", "text/plain")

		response.Write([]byte(url))

	} else {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Only post method type"))
	}
}

func redirectURLHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		body := ""

		for key, value := range request.URL.Query() {
			body += fmt.Sprintf("%s: %v\r\n", key, value)
		}

		response.WriteHeader(http.StatusTemporaryRedirect)
		response.Header().Set("Location", request.URL.Path)

		response.Write([]byte(body))
	} else {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Only get method type"))
	}
}
