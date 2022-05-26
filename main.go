package main

import (
	"fmt"
	"log"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	port := ":8085"
	host := fmt.Sprintf("localhost%s", port)
	// create a new server
	mux := http.NewServeMux()
	//handlers
	mux.HandleFunc("/pings", ChainMiddlewares(ping))
	mux.HandleFunc("/ping", ChainMiddlewares(ping, LoggerMiddleware))
	// start server
	log.Printf("Listening on %s", port)
	log.Fatal(http.ListenAndServe(host, mux))
}
