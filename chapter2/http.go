package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Home struct {
	Message string `json:"message"`
}

func (home Home) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	errorEncoder := json.NewEncoder(response).Encode(home)

	if request.URL.Path != "/home" {
		http.NotFound(response, request)
	}

	if errorEncoder != nil {
		response.WriteHeader(http.StatusInternalServerError)
	}
}

func FindCepHandler(response http.ResponseWriter, request *http.Request) {

	cep := request.URL.Query().Get("cep")

	viaCep, errorFindCep := FindCep(cep)

	if errorFindCep != nil {
		response.WriteHeader(http.StatusInternalServerError)
	}

	response.Header().Set("Content-Type", "application/json")
	errorEncoder := json.NewEncoder(response).Encode(viaCep)

	if errorEncoder != nil {
		response.WriteHeader(http.StatusInternalServerError)
	}

	response.WriteHeader(http.StatusOK)
}

func FindCep(cep string) (*ViaCep, error) {

	var url string = "https://viacep.com.br/ws/" + cep + "/json/"

	response, errorRequest := http.Get(url)

	if errorRequest != nil {
		return nil, errorRequest
	}

	data, errorRead := io.ReadAll(response.Body)

	if errorRead != nil {
		return nil, errorRead
	}

	var viaCep ViaCep

	errorUnmarshal := json.Unmarshal(data, &viaCep)

	if errorUnmarshal != nil {
		return nil, errorUnmarshal
	}

	return &viaCep, nil
}

func Lesson5() {

	fmt.Println("Lesson 5 - Http")

	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/address", FindCepHandler)
	serverMux.Handle("/", Home{Message: "Welcome to the Home"})
	errorServerHttp := http.ListenAndServe(":8080", serverMux)

	if errorServerHttp != nil {
		panic(errorServerHttp)
	}
}
