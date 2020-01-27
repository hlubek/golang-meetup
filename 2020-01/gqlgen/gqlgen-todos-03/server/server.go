package main

import (
	"log"
	"net/http"
	"os"

	gqlgen_todos "example/gqlgen-todos"

	"github.com/99designs/gqlgen/handler"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	todosRepo, err := gqlgen_todos.NewTodosRepository("data.json")
	if err != nil {
		log.Fatalf("error opening todos repository: %v", err)
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query",
		handler.GraphQL(
			gqlgen_todos.NewExecutableSchema(
				gqlgen_todos.Config{Resolvers: gqlgen_todos.NewResolver(todosRepo)},
			),
		),
	)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
