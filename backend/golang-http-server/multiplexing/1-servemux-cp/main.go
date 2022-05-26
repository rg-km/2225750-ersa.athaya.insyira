package main

import (
	"fmt"
	"net/http"
	"time"
)

// Dari contoh yang telah diberikan, buatlah http.ServeMux yang memiliki dua route
// Route pertama yaitu "/time" yang menghandle TimeHandler
// Route kedua yaitu "/hello" yang menghandle SayHelloHandler

func TimeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		fmt.Fprintf(w, "%v, %v %v %v", t.Weekday(), t.Day(), t.Month(), t.Year())
	} // TODO: replace this
}

func SayHelloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// tangkap query param name
		name := r.URL.Query().Get("name")
		if name != "" {
			fmt.Fprintf(w, "Hello, %s!", name)
		} else {
			fmt.Fprintf(w, "Hello there")
		}
	} // TODO: replace this
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	// TODO: answer here
	mux.HandleFunc("/time", TimeHandler())
	mux.HandleFunc("/hello", SayHelloHandler())
	return mux

}
