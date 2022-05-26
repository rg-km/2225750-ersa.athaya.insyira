package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Dari contoh yang telah diberikan, buatlah route untuk TimeHandler dan SayHelloHandler.
// Buatlah route "/time" pada TimeHandler dan "/hello" untuk SayHelloHandler dengan menggunakan ServeMux

var TimeHandler = func(writer http.ResponseWriter, request *http.Request) {
	// TODO: answer here
	if request.Method == "GET" {
		t := time.Now()
		fmt.Fprintf(writer, "%v, %v %v %v", t.Weekday(), t.Day(), t.Month(), t.Year())
		writer.WriteHeader(http.StatusOK)
	}

	// jika tidak menggunakan GET
	writer.WriteHeader(http.StatusMethodNotAllowed)
}

var SayHelloHandler = func(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
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

	// jika bukan GET
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func main() {
	mux := http.NewServeMux()
	// TODO: answer here
	mux.HandleFunc("/time", TimeHandler)
	mux.HandleFunc("/hello", SayHelloHandler)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
