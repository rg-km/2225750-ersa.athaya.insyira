package main

import "fmt"

func main() {
	//membuat rectangle dengan anonymous struct
	//field dari struct ini sama seperti rectangle sebelumnya
	// TODO: answer here

	fmt.Println("")
	fmt.Println("")

	newRectangle := struct {
		Width  int
		Length int
	}{
		Width:  10,
		Length: 30,
	}

	fmt.Printf("rectangle dengan lebar %d dan panjang %d", newRectangle.Width, newRectangle.Length)
	fmt.Println("")
	fmt.Println("")

	Mahasiswa := struct {
		Nama  string
		Nim   string
		Kelas string
	}{
		Nama:  "Rizki",
		Nim:   "A11.2018.09876",
		Kelas: "A11.2018",
	}

	fmt.Println("Nama : ", Mahasiswa.Nama)
	fmt.Println("Nim : ", Mahasiswa.Nim)
	fmt.Println("Kelas : ", Mahasiswa.Kelas)
	fmt.Println("")
	fmt.Println("")
}
