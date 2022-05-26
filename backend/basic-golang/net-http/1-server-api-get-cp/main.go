package main

import (
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
)

// Dari contoh sebelumnya tambahkan filter untuk
// menampilkan meja berdasarkan total meja
// key input dari client menggunakan `total`
// contuh URL: http://localhost:8080/table?total=2

type Table struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Color string `json:"color"`
	Total int    `json:"total"`
}

func TableHandler(w http.ResponseWriter, r *http.Request) {
	// set header content-type dengan json untuk menentukan response type
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {

		// ambil value dari key `total`
		total := r.FormValue("total")

		// convert string ke integer
		totalInt, _ := strconv.Atoi(total)

		// cek apakah total yang diinput user adalah angka atau bukan
		if totalInt > 0 {

			// filter data meja berdasarkan total meja
			var result []Table
			for _, table := range data {
				if table.Total == totalInt {
					result = append(result, table)
				}
			}

			// jika total tidak ada dalam data,
			// maka return response table not found
			if len(result) == 0 {
				http.Error(w, `{"status":"table not found"}`, http.StatusNotFound)
				return
			}

			// encode data ke dalam format string JSON
			resultJSON, err := json.Marshal(result)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write(resultJSON)
			return

		} else {
			// kirim pesan error bad request
			http.Error(w, "invalid total", http.StatusBadRequest)
			return
		}
	}

	// jika method yang diinputkan bukan GET
	http.Error(w, "", http.StatusBadRequest)

}

var data = []Table{
	{ID: "T001", Type: "Meja Lipat", Color: "Coklat", Total: 3},
	{ID: "T002", Type: "Meja Belajar", Color: "Hitam", Total: 4},
	{ID: "T003", Type: "Meja Makan", Color: "Coklat", Total: 9},
	{ID: "T004", Type: "Meja Hejau", Color: "Hijau", Total: 3},
}

func main() {
	http.HandleFunc("/table", TableHandler)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
