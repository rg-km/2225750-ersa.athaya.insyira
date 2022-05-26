package main

import (
	"errors"
	"fmt"
)

type Person struct {
	name string
	age  int
}

func main() {
	people := []Person{{name: "Bob", age: 21}, {name: "Sam", age: 28}, {name: "Ann", age: 21}, {name: "Joe", age: 22}, {name: "Ben", age: 28}}
	fmt.Println(AgeDistribution(people))
	fmt.Println(FilterByAge(people, 2))
}

func AgeDistribution(people []Person) map[int]int {
	ageDistribution := make(map[int]int)
	for _, person := range people {
		if _, ok := ageDistribution[person.age]; ok {
			ageDistribution[person.age]++
		} else {
			ageDistribution[person.age] = 1
		}
	}
	return ageDistribution // TODO: replace this
}

func FilterByAge(people []Person, age int) ([]Person, error) {
	var filteredPeople []Person
	for _, person := range people {
		if person.age == age {
			filteredPeople = append(filteredPeople, person)
		}
	}

	if len(filteredPeople) == 0 {
		return filteredPeople,errors.New("Age not found")
	}
	return filteredPeople, nil // TODO: replace this
}
