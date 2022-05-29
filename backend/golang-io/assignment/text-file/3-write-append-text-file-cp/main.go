package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("dummyCommit")
}

func AddString(fileName string, stringToAdd string) error {
	//membuka file
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
		return err
	}

	//menutup file sebelum fungsi selesai dijalankan
	defer file.Close()

	_, err = file.WriteString(stringToAdd)

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
		return err
	}

	return nil
}
