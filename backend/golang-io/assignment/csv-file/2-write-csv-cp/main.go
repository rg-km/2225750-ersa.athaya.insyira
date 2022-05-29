package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Menu struct {
	Name  string
	Price int
}

func WriteToCSV(fileName string, records []Menu) error {
	csvFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	// TODO: answer here
	// Write some records to file
	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	// Using Write
	for _, record := range records {
		row := []string{record.Name, strconv.Itoa(record.Price)}
		if err := csvWriter.Write(row); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	csvFileName := "menu.csv"
	data := []Menu{
		{"Pizza", 100},
		{"Burger", 200},
		{"Coffee", 300},
		{"Tea", 400},
		{"Sandwich", 500},
	}

	_ = WriteToCSV(csvFileName, data)
}
