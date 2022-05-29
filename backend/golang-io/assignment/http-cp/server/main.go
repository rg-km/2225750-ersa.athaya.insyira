package main

import (
	"net/http"
	"strconv"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "GET" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello, world!"))
			return
		}

	})
	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "GET" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(r.URL.Query().Get("data")))
			return
		}
	})
	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		w.Header().Set("Content-Type", "application/json")

		//request method is GET
		//returns the sum of a and b
		if r.Method == "GET" {
			sum := 0
			a, err := strconv.Atoi(r.URL.Query().Get("a"))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("a is not a number"))
				return
			}
			b, err := strconv.Atoi(r.URL.Query().Get("b"))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("b is not a number"))
				return
			}
			sum = a + b

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(strconv.Itoa(sum)))
			return
		}
	})
	mux.HandleFunc("/hellojson", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "GET" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "Hello, world!"}`))
			return
		}

	})

	return mux
}
func main() {
	http.ListenAndServe(":8080", Routes())
}
