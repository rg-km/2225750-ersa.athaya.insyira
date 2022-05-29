package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fileName := "questions.csv"
	data := make(map[string]string)

	data, err := CSVToMap(data, fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(data)
	fmt.Print("dummy commit")

}

func CSVToMap(data map[string]string, fileName string) (map[string]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.LazyQuotes = true

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		data[record[0]] = record[1]
	}

	return data, nil
}
