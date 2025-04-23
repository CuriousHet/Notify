package main

import (
	"log"
	"net/http"

	"github.com/CuriousHet/Notify/gql"
	"github.com/CuriousHet/Notify/notification"
	"github.com/CuriousHet/Notify/server"
	"github.com/gorilla/mux"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
	queue := notification.NewQueue(100)
	dispatcher := notification.NewDispatcher(queue, 3)
	dispatcher.Start(5) // Start 5 workers

	go server.StartGRPCServer(queue)

	schema := graphql.MustParseSchema(gql.SchemaString, &gql.Resolver{})

	// Set up HTTP routes (GraphQL queries)
	r := mux.NewRouter()
	r.Handle("/query", &relay.Handler{Schema: schema}) // This is the correct way to handle GraphQL queries

	// Start HTTP server for GraphQL (port 8081)
	log.Println("GraphQL server is running on :8081...")
	http.ListenAndServe(":8081", r)
}
