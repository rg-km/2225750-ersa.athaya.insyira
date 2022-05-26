package main

import (
	"bytes"
	"fmt"
	"html/template"
)

// Dari contoh yang telah diberikan, cobalah untuk mempraktikkan data evaluation pada template.
// Lengkapi function ExecuteToByteBuffer sehingga objek dari struct Account dapat ter-render ke dalam template.
// Gunakan bytes.Buffer sebagai io.Writer pada template.
// Lengkapi juga textTemplate, sehingga variabel dari objek Account dapat ter-render ke dalam template.
// Contoh:
// acc := {Name: "Tony", Number: "1002321", Balance: 1000}
// Akun Tony dengan Nomor 1002321 memiliki saldo sebesar $1000.

type Account struct {
	Name    string
	Number  string
	Balance int
}

func ExecuteToByteBuffer(account Account) ([]byte, error) {
	// siapkan kalimat yang akan dicetak
	textTemplate := "Akun {{.Name}} dengan Nomor {{.Number}} memiliki saldo sebesar ${{.Balance}}."

	// membuat tempalte dengan nama account
	t, err := template.New("account").Parse(textTemplate)
	if err != nil {
		return nil, err
	}

	// membuat buffer
	var b bytes.Buffer

	// menjalankan template
	err = t.Execute(&b, account)
	if err != nil {
		return nil, err
	}

	// mengembalikan hasil
	return b.Bytes(), nil

}

func main() {
	// isi struct account
	acc := Account{
		Name:    "Tony",
		Number:  "1002321",
		Balance: 1000,
	}
	// memanggil function ExecuteToByteBuffer
	b, err := ExecuteToByteBuffer(acc)

	// cek error
	if err != nil {
		panic(err)
	}
	// menampilkan hasil dengan convert ke string
	fmt.Println(string(b))
}
