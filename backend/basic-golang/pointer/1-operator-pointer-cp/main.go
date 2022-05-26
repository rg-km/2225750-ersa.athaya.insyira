package main

import "fmt"

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan address operator dan indirect operator.

// Print alamat memori dari masing-masing variabel dibawah ini
func main() {
	name := "Roger"
	age := 20
	isMarried := true

	// TODO: answer here
	fmt.Printf("Name : %v\n", &name)
	fmt.Printf("Age : %v\n", &age)
	fmt.Printf("IsMarried : %v\n", &isMarried)
}
