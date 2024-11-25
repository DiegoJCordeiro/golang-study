package main

import (
	"context"
	"database/sql"
	"github.com/DiegoJCordeiro/golang-study/activity/server/cfg"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"
)

func main() {

	configuration, err := cfg.LoadConfiguration("./cmd/server")

	if err != nil {
		panic(err)
	}

	_, errDB := sql.Open(configuration.DBDriver, configuration.DBHost)

	if errDB != nil {
		panic(errDB)
	}

	var router chi.Router = chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(Recover)

	router.Get("/cotacao", func(w http.ResponseWriter, r *http.Request) {

		client := http.Client{}
		ctx, errCtx := context.WithTimeout(context.Background(), 200*time.Millisecond)

		if errCtx != nil {
			panic(errCtx)
		}

		request, errorReq := http.NewRequestWithContext(
			ctx,
			"GET",
			"https://economia.awesomeapi.com.br/json/last/USD-BRL",
			nil,
		)

		if errorReq != nil {
			panic(errorReq)
		}

		_, errorResp := client.Do(request)

		if errorResp != nil {
			panic(errorResp)
		}

	})

	_ = http.ListenAndServe(":8080", router)
}

func Recover(next http.Handler) http.Handler {
	recoverHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("Panic: %v \n", rec)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})

	return http.HandlerFunc(recoverHandler)
}
