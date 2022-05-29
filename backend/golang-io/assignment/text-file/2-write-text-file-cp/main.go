package main

import (
	"fmt"
	"log"
	"os"
)

// dalam test ini terdapat fungsi os.Remove ya. itu automatis nge remove file yang telah dibuat
// Untuk keperluan testing
func WriteFile(fileName string, fileData string) error {

	// membuat file
	file, err := os.Create(fileName)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	// menutup file sebelum fungsi selesai dijalankan
	defer file.Close()

	len, err := file.WriteString(fileData)

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	fmt.Printf("File Name: %s\n", file.Name())
	fmt.Printf("Length: %d bytes\n", len)

	return nil
}
