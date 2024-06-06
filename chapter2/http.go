package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Mail struct {
	Message string `json:"message"`
}

func (home Mail) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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

	var url = "https://viacep.com.br/ws/" + cep + "/json/"

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

func TemplateHandler(response http.ResponseWriter, request *http.Request) {

	var errorTemplate error

	templates := []string{
		"files/template-mail.html",
		"files/header-mail.html",
		"files/footer-mail.html",
	}

	parseBool, errorTemplate := strconv.ParseBool(request.URL.Query().Get("single"))

	if parseBool == true {

		templateMailsHTML := template.Must(template.New("template-mail").Parse("<h1>The message of Mail is: {{ .Message }}</h1>"))
		errorTemplate = templateMailsHTML.Execute(response, Mail{Message: "Message Mail."})
		response.Header().Add("Content-Type", "text/html")
	} else if parseBool == false {

		templateMailsHTML := template.Must(template.New("template-mail.html").Funcs(template.FuncMap{"ToUpperCase": ToUpperCase}).ParseFiles(templates...))
		errorTemplate = templateMailsHTML.Execute(response, []Mail{{Message: "1 - Message Mail."}, {Message: "2 - Message Mail."}})
		response.Header().Add("Content-Type", "text/html")
	} else {
		response.WriteHeader(http.StatusNotFound)
	}

	if errorTemplate != nil {
		response.WriteHeader(http.StatusInternalServerError)
	}
}

func ToUpperCase(text string) string {

	return strings.ToUpper(text)
}

func Lesson5() {

	fmt.Println("Lesson 5 - Http")

	serverMux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))
	serverMux.Handle("/", fileServer)
	serverMux.Handle("/mail", Mail{Message: "Welcome to the Mail"})

	serverMux.HandleFunc("/address", FindCepHandler)
	serverMux.HandleFunc("/handler-template", TemplateHandler)

	errorServerHttp := http.ListenAndServe(":8080", serverMux)

	if errorServerHttp != nil {
		panic(errorServerHttp)
	}
}
