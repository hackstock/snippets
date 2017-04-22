package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		//NOTE : Invoke ParseForm or ParseMultipartForm before reading form values
		r.ParseForm()
		/*
			Loops through r.Form object to print key-value pairs of
			url-form encoded data. Note that these include both data sent
			through request url and request body
		*/
		for key, value := range r.Form {
			fmt.Printf("%s = %s\n", key, value)
		}

		fmt.Fprintln(w, "Registeration successful. You inner man is now aligned with nature")
	})

	http.ListenAndServe(":8080", mux)
}
