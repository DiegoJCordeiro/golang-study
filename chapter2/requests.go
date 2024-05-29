package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func requestHttpGet() {

	httpGet, errorRequest := http.Get("https://www.google.com/")

	if errorRequest != nil {
		panic(errorRequest)
	}

	response, errReader := io.ReadAll(httpGet.Body)

	if errReader != nil {
		panic(errorRequest)
	}

	defer writeResponseToFile(response)

	errHttpClose := httpGet.Body.Close()

	if errHttpClose != nil {
		panic(errHttpClose)
	}
}

func requestHttpPost() {

	httpPost, errorRequest := http.Post("https://buscacepinter.correios.com.br/app/endereco/carrega-cep-endereco.php", "application/json", io.Reader(nil))

	if errorRequest != nil {
		panic(errorRequest)
	}

	response, errReader := io.ReadAll(httpPost.Body)

	if errReader != nil {
		panic(errorRequest)
	}

	defer writeResponseToFile(response)

	errHttpClose := httpPost.Body.Close()

	if errHttpClose != nil {
		panic(errHttpClose)
	}
}

func writeResponseToFile(response []byte) {

	file, errorOpenFile := os.OpenFile("./files/requests.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if errorOpenFile != nil {
		panic(errorOpenFile)
	}

	writer := bufio.NewWriter(file)

	_, errWrite = writer.WriteString("\n")
	_, errWrite = writer.Write([]byte(response))

	errFlush := writer.Flush()
	errFileClose := file.Close()

	if errFileClose != nil {
		panic(errFileClose)
	}

	if errFlush != nil {
		panic(errFlush)
	}
}

func Lesson2() {

	fmt.Println("Lesson 2 - Http Requests")

	requestHttpGet()
	requestHttpPost()

	fmt.Println("---")
}
