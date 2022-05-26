package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

// Buat view untuk menampilkan table berdasarkan table ID
// URL: http://localhost:8080/view/table?id=T001

type Table struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Color string `json:"color"`
	Total int    `json:"total"`
}

func TablesHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	// set header
	w.Header().Set("Content-Type", "text/html")

	// ambil id yang diinput user
	id := r.URL.Query().Get("id")

	// cek ke data apakah ada id yang di minta ?
	for _, table := range data {
		if table.ID == id {

			// masukan ke dalam map
			data := map[string]interface{}{
				"title": "Server-Web-API",
				"table": table,
			}

			// untuk menambahkan delimiter path
			filepath := path.Join("views", "index.html")
			tmpl, err := template.ParseFiles(filepath)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			// render template
			if err := tmpl.Execute(w, data); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}

	// jika tidak ada, tampilkan pesan
	fmt.Fprintf(w, "Table ID: %s not found", id)

}

var data = []Table{
	{ID: "T001", Type: "Meja Lipat", Color: "Coklat", Total: 3},
	{ID: "T002", Type: "Meja Belajar", Color: "Hitam", Total: 4},
	{ID: "T003", Type: "Meja Makan", Color: "Coklat", Total: 9},
	{ID: "T004", Type: "Meja Hejau", Color: "Hijau", Total: 3},
}

func main() {
	http.HandleFunc("/view/table", TablesHandler)

	fmt.Println("starting web server at http://localhost:8080/")
	// daftarkan server untuk listen di port 8080
	http.ListenAndServe(":8080", nil)
}
