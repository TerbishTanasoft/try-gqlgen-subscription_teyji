// package main

// import (
// 	"log"
// 	"net/http"
// 	"os"

// 	"try-gqlgen-subscription_teyji/graph"

// 	"github.com/99designs/gqlgen/graphql/handler"
// 	"github.com/99designs/gqlgen/graphql/handler/transport"
// 	"github.com/99designs/gqlgen/graphql/playground"
// )

// const defaultPort = "8080"

// func main() {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = defaultPort
// 	}

// 	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

// 	srv.AddTransport(&transport.Websocket{}) // <---- This is the important part!

// 	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
// 	http.Handle("/query", srv)

//		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
//		log.Fatal(http.ListenAndServe(":"+port, nil))
//	}
//
// main.go
package main

import (
	"log"
	"net/http"
	"os"

	"try-gqlgen-subscription_teyji/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", cors.Default().Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
