package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	ctx    context.Context
	cancel context.CancelFunc
)

func init() {

	ctx = context.Background()
	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	ctx = context.WithValue(ctx, "token", "asd675as5d6a5d12325678.dad81273.dwr1298")
}

func requestHttpWithContext() {

	request, errorRequest := http.NewRequestWithContext(ctx, "GET", "https://jsonplaceholder.typicode.com/posts", nil)

	if errorRequest != nil {
		panic(errorRequest)
	}

	response, errorResponse := http.DefaultClient.Do(request)

	if errorResponse != nil {
		panic(errorResponse)
	}

	responseRead, errorReadAll := io.ReadAll(response.Body)

	if errorReadAll != nil {
		panic(errorReadAll)
	}

	fmt.Printf("%s \n", string(responseRead))
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)
}

func recoveryContextValue() {
	token := ctx.Value("token")
	fmt.Printf("Value Token is: %s\n", token)
}

func cancelContext() {
	cancel()
}

func Lesson1() {

	fmt.Printf("Lesson 1 - Contexts\n")
	recoveryContextValue()
	requestHttpWithContext()
	defer cancelContext()
}
