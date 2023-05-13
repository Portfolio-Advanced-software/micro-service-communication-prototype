package main

import (
	"log"
	"net/http"

	"github.com/Portfolio-Advanced-software/micro-service-communication-prototype/GraphQL/schema"
	"github.com/gorilla/mux"
	"github.com/graphql-go/handler"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		h := handler.New(&handler.Config{
			Schema:   &schema.Schema,
			Pretty:   true,
			GraphiQL: true,
		})
		h.ServeHTTP(w, r)
	})

	log.Println("Server started on http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", r))
}
