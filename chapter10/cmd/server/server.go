package main

import (
	"database/sql"
	"github.com/DiegoJCordeiro/golang-study/chapter10/internal/infra/database"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/DiegoJCordeiro/golang-study/chapter10/graph"
)

const defaultPort = "8080"

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:1234@localhost:5432/postgres?sslmode=disable")

	if err != nil {
		log.Fatalf("failed to open connection with database: %v", err)
	}

	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	courseDB := database.NewCourseDB(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CategoryDB: categoryDB,
		CourseDB:   courseDB,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
