package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Dari contoh yang diberikan, cobalah untuk mengimplementasikan sebuah http client sederhana.
// Buatlah sebuah http client dan lakukan request GET ke API https://www.metaweather.com/api/.
// Buatlah request get ke endpoint /api/location/(woeid)/(date)/ dengan nilai woeid 1047378.
// Untuk date, gunakan format YYYY/MM/dd

func main() {
	// TODO: answer here
	
	// inisialisasi client http
	myClient := &http.Client{}

	// lakukan hit endpoint
	resp, err := myClient.Get("https://www.metaweather.com/api/location/1047378/2019/10/10")
	// cek error
	if err != nil {
		log.Fatal(err)
	}
	// defer
	defer resp.Body.Close()
	// ubah resp menjadi byte dengan membaca semua respon body
	body, err := ioutil.ReadAll(resp.Body)
	// cek error
	if err != nil {
		log.Fatal(err)
	}
	// print respon body
	log.Println(string(body))

}
