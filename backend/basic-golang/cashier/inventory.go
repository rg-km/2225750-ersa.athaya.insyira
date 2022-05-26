package main

import (
	"fmt"
)

type Item struct {
	id    int
	name  string
	price float64
}

type Inventory struct {
	items []Item
}

func NewItem(id int, name string, price float64) Item {
	return Item{id, name, price}
}

func (i *Inventory) Add(item Item) {
	// TODO: answer here
	// tambahkan item ke inventory
	i.items = append(i.items, item)

}

func (i Inventory) Items() []Item {
	return i.items
}

func main() {
	// inisiasi struct Inventory
	inventory := Inventory{}

	// tambahkan item ke inventory
	inventory.Add(NewItem(1, "item1", 10.0))
	inventory.Add(NewItem(2, "item2", 20.0))

	// tampilkan semua item
	for _, item := range inventory.Items() {
		fmt.Println(item)
	}

}
