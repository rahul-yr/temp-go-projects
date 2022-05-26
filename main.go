package main

import (
	"fmt"
	"log"
	"net/http"
)

var commonMiddlewares = []CustomMiddlewares{
	ApiMethodvalidatorMiddleware,
	LoggerMiddleware,
}

func main() {
	port := ":8085"
	host := fmt.Sprintf("localhost%s", port)
	// create a new server
	mux := http.NewServeMux()
	// register routes
	for route, handler := range AvailableRoutes {
		mux.HandleFunc(route, ChainMiddlewares(handler.HandlerFunc, commonMiddlewares...))
	}

	// start server
	log.Printf("Listening on %s", port)
	log.Fatal(http.ListenAndServe(host, mux))
}
