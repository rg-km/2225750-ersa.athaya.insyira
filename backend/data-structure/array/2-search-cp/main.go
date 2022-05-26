package main

import "fmt"

func main() {
	var cars1 = []string{"Toyota", "Honda", "Nissan", "BMW", "Chevy", "Ford"}
	var cars2 = []string{"Ford", "BMW", "Audi", "Mercedes"}

	result, err := SearchMatch(cars1, cars2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Matched car brand: ", result)
	}
}

func SearchMatch(arr1 []string, arr2 []string) ([]string, error) {
	var result []string
	for _, v1 := range arr1 {
		for _, v2 := range arr2 {
			if v1 == v2 {
				result = append(result, v1)
			}
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no match")
	}
	return result, nil
}
