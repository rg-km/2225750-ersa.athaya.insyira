package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Dari contoh sebelumnya tambahkan logic
// untuk handle POST request add bulk table
// endpoint URL: http://localhost:8080/tables
// request body:
/*
[
    {
        "id": "T012",
        "type": "Meja Rias",
        "color": "Coklat",
        "total": 3
    },
    {
        "id": "T011",
        "type": "Meja Kotak",
        "color": "Hitam",
        "total": 4
    }
]
*/

type Table struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Color string `json:"color"`
	Total int    `json:"total"`
}

func TablesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// validate methode request
	// logic handler untuk GET request
	if r.Method == "GET" {
		// encode data ke dalam format string JSON
		result, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// untuk mendaftarkan result sebagai response
		w.Write(result)
		return
	}

	// logic handle POST request
	if r.Method == "POST" {
		// TODO: answer here
		// inisialisasi variable untuk menampung data dari request body
		// pakai slice karena yang mau diinput banyak data sekaligus
		var tables []Table

		// baca request body dari client
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// unmarshal data dari request body ke dalam variable `tables`
		// jika gagal, kembalikan response error
		if err := json.Unmarshal(body, &tables); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// tambah data hasil unmarshal ke dalam variable `data`
		data = append(data, tables...)

		// set header response code with status created/201
		w.WriteHeader(http.StatusCreated)
		// write json reponse body
		w.Write([]byte(`{"status":"add tables succeed"}`))
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func showData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		result, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
	}

	http.Error(w, "", http.StatusBadRequest)
}

var data = []Table{
	{ID: "T001", Type: "Meja Lipat", Color: "Coklat", Total: 3},
	{ID: "T002", Type: "Meja Belajar", Color: "Hitam", Total: 4},
	{ID: "T003", Type: "Meja Makan", Color: "Coklat", Total: 9},
	{ID: "T004", Type: "Meja Hejau", Color: "Hijau", Total: 3},
}

func main() {
	http.HandleFunc("/tables", TablesHandler)
	http.HandleFunc("/table", showData)

	fmt.Println("starting web server at http://localhost:8080/")
	// daftarkan server untuk listen di port 8080
	http.ListenAndServe(":8080", nil)
}
