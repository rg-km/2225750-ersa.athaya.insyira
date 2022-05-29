package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	res, _ := ReadFile("./read.txt")
	fmt.Printf("File Name: %s", res.Name)
	fmt.Printf("\nSize: %d bytes", res.Size)
	fmt.Printf("\nData: %s\n", res.Data)

	//fmt.Println("dummyCommit")
}

// Gunakan struct untuk menyimpan data file nya ya
type FileData struct {
	Name string
	Size int
	Data []byte
}

func ReadFile(name string) (FileData, error) {
	// nama text file yang ingin dibaca
	fileName := name

	//membaca text file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return FileData{}, err
	}
	return FileData{
		Name: fileName,
		Size: len(data),
		Data: data,
	}, nil
}
