package main

import (
	"log"
	"net/http"
	"time"
)

type CustomMiddlewares func(http.HandlerFunc) http.HandlerFunc

func ChainMiddlewares(next http.HandlerFunc, middlewares ...CustomMiddlewares) http.HandlerFunc {
	// iterate the middlewares from last to first
	// to retain the order of middlewares
	for i := len(middlewares) - 1; i >= 0; i-- {
		next = middlewares[i](next)
	}

	return next
}

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// log the request
		log.Printf("Request[%s %s %s]", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
		end := time.Now()
		// log the response
		log.Printf("Response[%s %s %s %s]", r.RemoteAddr, r.Method, r.URL, end.Sub(start))
	})
}

func ApiMethodvalidatorMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		registeredRoute := AvailableRoutes[r.URL.Path]
		if registeredRoute.Method != r.Method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
