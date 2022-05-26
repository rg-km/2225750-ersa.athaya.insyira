package main

import "encoding/json"

// Dari contoh sebelumnya
// buat JSON string array nested seperti berikut
/*
[
{
		"jenis": "Meja Makan",
		"warna": "Coklat",
		"jumlah": 20,
		"ukuran": {
				"panjang": "50 cm",
				"tinggi": "25 cm"
		}
},
{
		"jenis": "Meja Lipat",
		"warna": "Hitam",
		"jumlah": 1,
		"ukuran": {
				"panjang": "70 cm",
				"tinggi": "30 cm"
		}
}
]
*/

type Ukuran struct {
	// TODO: answer here
	Panjang string `json:"panjang"`
	Tinggi  string `json:"tinggi"`
}

type Meja struct {
	// TODO: answer here
	Jenis  string `json:"jenis"`
	Warna  string `json:"warna"`
	Jumlah int    `json:"jumlah"`
	Ukuran Ukuran `json:"ukuran"`
}

type Items struct {
	MejaMeja []Meja
}

func (m Items) EncodeJSON() string {
	// TODO: answer here
	// encode
	data, _ := json.Marshal(m.MejaMeja)
	// convert to string
	return string(data)
}

func NewMeja(m Items) Items {
	return m
}

func main() {

	// define ukuran
	ukuran1 := Ukuran{
		Panjang: "50 cm",
		Tinggi:  "25 cm",
	}

	ukuran2 := Ukuran{
		Panjang: "70 cm",
		Tinggi:  "30 cm",
	}

	// kita isi nilai ke struct Items
	item := Items{
		MejaMeja: []Meja{
			{
				Jenis:  "Meja Makan",
				Warna:  "Coklat",
				Jumlah: 20,
				Ukuran: ukuran1,
			},
			{
				Jenis:  "Meja Lipat",
				Warna:  "Hitam",
				Jumlah: 1,
				Ukuran: ukuran2,
			},
		},
	}

	// panggail newMeja
	meja := NewMeja(item)

	// encode JSON
	result := meja.EncodeJSON()

	// cetak hasil
	println(result)
}
