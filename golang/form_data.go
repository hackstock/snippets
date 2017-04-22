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
			Reads individual key-value pairs from
			r.Form object. Note that these include both data sent
			through request url and request body
		*/
		fmt.Printf("USERNAME => %s\n", r.FormValue("username"))
		fmt.Printf("EMAIL => %s\n", r.FormValue("email"))
		fmt.Printf("PASSWORD => %s\n", r.FormValue("password"))
		fmt.Printf("PAGEINDEX => %s\n", r.FormValue("pageindex"))

		fmt.Fprintln(w, "Registeration successful. You inner man is now aligned with nature")
	})

	http.ListenAndServe(":8080", mux)
}
