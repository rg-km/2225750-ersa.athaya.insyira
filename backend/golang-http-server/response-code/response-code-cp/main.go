package main

import (
	"net/http"
)

var names = []string{
	"Tony",
	"Roger",
	"Bruce",
	"Stephen",
}

func IsNameExists(name string) bool {
	for _, n := range names {
		if n == name {
			return true
		}
	}

	return false
}

func GetNameHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		if r.Method == "GET" {
			// ambil param name
			name := r.URL.Query().Get("name")
			if !IsNameExists(name) {
				w.WriteHeader(http.StatusNotFound)
			} else if name == "" {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		} else {
			// jika method bukan GET
			w.WriteHeader(http.StatusMethodNotAllowed)
		}

	}
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	// TODO: answer here
	mux.HandleFunc("/name", GetNameHandler())
	return mux
}
