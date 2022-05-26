package main

import (
	"fmt"
	"net/http"
	"time"
)

// Dari contoh yang telah diberikan, buatlah route untuk TimeHandler dan SayHelloHandler.
// Buatlah route "/time" pada TimeHandler dan "/hello" untuk SayHelloHandler dengan menggunakan http.HandleFunc

var TimeHandler = func(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		t := time.Now()
		fmt.Fprintf(writer, "%v, %v %v %v", t.Weekday(), t.Day(), t.Month(), t.Year())
		writer.WriteHeader(http.StatusOK)
	}

	// jika tidak menggunakan GET
	writer.WriteHeader(http.StatusMethodNotAllowed)

	// TODO: answer here
}

var SayHelloHandler = func(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// ambil query param name
		name := r.URL.Query().Get("name")
		if name == "" {
			fmt.Fprintf(w, "Hello there")
			w.WriteHeader(http.StatusOK)
		} else {
			fmt.Fprintf(w, "Hello, %s!", name)
			w.WriteHeader(http.StatusOK)
		}
	}

	// jika method bukan GET
	w.WriteHeader(http.StatusMethodNotAllowed)
	// TODO: answer here
}

func main() {
	http.HandleFunc("/time", TimeHandler)
	http.HandleFunc("/hello", SayHelloHandler)
	http.ListenAndServe("localhost:8080", nil)
}
