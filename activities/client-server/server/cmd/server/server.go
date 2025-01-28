package main

import (
	"database/sql"
	"encoding/json"
	"github.com/DiegoJCordeiro/golang-study/activity/server/cfg"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/dto"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"time"
)

func main() {

	configuration, err := cfg.LoadConfiguration("./cmd/server")
	if err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}

	db, err := sql.Open(configuration.DBDriver, configuration.DBHost)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Erro ao verificar conexão com o banco de dados: %v", err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Get("/cotacao", HandlerQuotation)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Servidor iniciado na porta 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Erro no servidor HTTP: %v", err)
	}
}

func HandlerQuotation(w http.ResponseWriter, r *http.Request) {

	const apiURL = "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	ctx := r.Context()

	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		http.Error(w, "Erro ao criar requisição para API de câmbio", http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Erro ao obter dados da API de câmbio", http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "API de câmbio retornou erro", http.StatusBadGateway)
		return
	}

	var quotation dto.QuotationDTO

	if err := json.NewDecoder(resp.Body).Decode(&quotation); err != nil {
		http.Error(w, "Erro ao decodificar resposta da API", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(quotation); err != nil {
		http.Error(w, "Erro ao enviar resposta JSON", http.StatusInternalServerError)
		return
	}
}
