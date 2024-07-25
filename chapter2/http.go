package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Posts struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

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

func SendPostsHandler(response http.ResponseWriter, request *http.Request) {

	var posts Posts
	requestBody, errorReadAll := io.ReadAll(request.Body)

	if errorReadAll != nil {
		panic(errorReadAll)
	}

	errorJsonUnmarshal := json.Unmarshal(requestBody, &posts)

	if errorJsonUnmarshal != nil {
		panic(errorJsonUnmarshal)
	}

	posts, errorPosts := sendPosts(&posts)

	if errorPosts != nil {
		response.WriteHeader(http.StatusInternalServerError)
	}

	response.Header().Set("Content-Type", "application/json")
	errorEncoder := json.NewEncoder(response).Encode(posts)

	if errorEncoder != nil {
		response.WriteHeader(http.StatusInternalServerError)
	}
}

func sendPosts(posts *Posts) (Posts, error) {

	var postsResponse Posts
	postsMarshaled, errorMarshal := json.Marshal(posts)

	if errorMarshal != nil {
		return postsResponse, errorMarshal
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)

	defer cancel()

	requestPosts, errorPost := http.NewRequestWithContext(ctx, "POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(postsMarshaled))

	if errorPost != nil {
		return postsResponse, errorPost
	}

	responsePosts, errorDefaultClient := http.DefaultClient.Do(requestPosts)

	if errorDefaultClient != nil {
		return postsResponse, errorDefaultClient
	}

	responseRead, errorReadAll := io.ReadAll(responsePosts.Body)

	defer responsePosts.Body.Close()

	if errorReadAll != nil {
		return postsResponse, errorReadAll
	}

	errorUnMarshal := json.Unmarshal(responseRead, &postsResponse)

	if errorUnMarshal != nil {
		return postsResponse, errorUnMarshal
	}

	return postsResponse, nil
}

func FindCepHandler(response http.ResponseWriter, request *http.Request) {

	cep := request.URL.Query().Get("cep")

	viaCep, errorFindCep := FindCep(cep)

	if errorFindCep != nil {
		response.WriteHeader(http.StatusInternalServerError)
	}

	response.WriteHeader(http.StatusOK)
	response.Header().Set("Content-Type", "application/json")
	viaCep.HttpMethod = "GET"
	errorEncoder := json.NewEncoder(response).Encode(viaCep)

	if errorEncoder != nil {
		response.WriteHeader(http.StatusInternalServerError)
	}
}

func FindCep(cep string) (*ViaCep, error) {

	var url = "https://viacep.com.br/ws/" + cep + "/json/"

	httpClient := http.Client{Timeout: 5 * time.Second}
	response, errorRequest := httpClient.Get(url)

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

	viaCep.HttpMethod = "GET"

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

	serverMux.Handle("GET /", fileServer)
	serverMux.Handle("GET /mail", Mail{Message: "Welcome to the Mail"})

	serverMux.HandleFunc("GET /address", FindCepHandler)
	serverMux.HandleFunc("GET /handler-template", TemplateHandler)
	serverMux.HandleFunc("POST /posts", SendPostsHandler)

	errorServerHttp := http.ListenAndServe(":8080", serverMux)

	if errorServerHttp != nil {
		panic(errorServerHttp)
	}
}
