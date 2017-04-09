package main

import (
	"fmt"
	"net/http"
)

func indexPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the home page")
}

func aboutPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the about page")
}

func loggerMiddleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request received : %v\n", r)
		nextHandler.ServeHTTP(w, r)
		fmt.Println("Request handled successfully")
	})
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", loggerMiddleware(http.HandlerFunc(indexPageHandler)))
	mux.Handle("/about", loggerMiddleware(http.HandlerFunc(aboutPageHandler)))

	http.ListenAndServe(":8080", mux)
}
