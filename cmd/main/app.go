package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Subj struct {
	Product string `json:"name"`
	Price   int    `json:"price"`
}

func JSONHandler(w http.ResponseWriter, req *http.Request) {
	// собираем данные
	subj := Subj{"Milk", 50}
	// кодируем в JSON
	resp, err := json.Marshal(subj)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// устанавливаем заголовок Content-Type
	// для передачи клиенту информации, кодированной в JSON
	w.Header().Set("content-type", "application/json")
	// устанавливаем код 200
	w.WriteHeader(http.StatusOK)
	// пишем тело ответа
	w.Write(resp)
}

func apiPage(response http.ResponseWriter, request *http.Request) {
	var greetingMessage = []byte("Api Page")
	response.Write(greetingMessage)
}

func mainPage(response http.ResponseWriter, request *http.Request) {
	var greetingMessage = []byte("Main Page")
	response.Write(greetingMessage)
}

func helperApi(response http.ResponseWriter, request *http.Request) {
	body := fmt.Sprintf("Method: %s\r\n", request.Method)

	//body += "Header ===============\r\n"
	//for k, v := range req.Header {
	//	body += fmt.Sprintf("%s: %v\r\n", k, v)
	//}
	//body += "Query parameters ===============\r\n"
	//for k, v := range request.URL.Query() {
	//	body += fmt.Sprintf("%s: %v\r\n", k, v)
	//}
	err := request.ParseForm()

	if err != nil {
		response.Write([]byte(err.Error()))
		return
	}

	for k, v := range request.Form {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}

	fmt.Print(response, body)
}

func reader(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "./app.go")
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(".."))

	mux.Handle(`/golang/`, http.StripPrefix(`/golang/`, fs))

	mux.HandleFunc("/", mainPage)
	mux.HandleFunc("/api", apiPage)
	mux.HandleFunc("/helper", helperApi)
	mux.HandleFunc("/reader", reader)

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		panic(err)
	}

}
