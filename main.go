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
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Initialize Prometheus metrics
	notification.InitMetrics()

	// Set up the notification queue and dispatcher
	queue := notification.NewQueue(100)
	dispatcher := notification.NewDispatcher(queue, 3)
	dispatcher.Start(5) // Start 5 worker goroutines

	// Start gRPC server in a separate goroutine
	go server.StartGRPCServer(queue)

	// Initialize GraphQL schema
	schema := graphql.MustParseSchema(gql.SchemaString, &gql.Resolver{})

	// Set up GraphQL server router
	r := mux.NewRouter()
	r.Handle("/query", &relay.Handler{Schema: schema}) // GraphQL endpoint
	r.Handle("/metrics", promhttp.HandlerFor(notification.CustomRegistry, promhttp.HandlerOpts{}))

	// Start HTTP server for GraphQL and metrics on port 8081
	log.Println("GraphQL and Metrics server running at http://localhost:8081...")
	http.ListenAndServe(":8081", r)
}
