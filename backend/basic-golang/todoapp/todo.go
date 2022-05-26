package main

import (
	"fmt"
	"time"
)

type Item struct {
	Title    string
	Deadline time.Time
}

type Todos struct {
	items []Item
}

func (todos *Todos) Add(item Item) {
	// TODO: answer here
	todos.items = append(todos.items, item)
}

func (todos *Todos) GetAll() []Item {
	// return []Item{} // TODO: replace this
	var todo []Item
	for _, item := range todos.items {
		todo = append(todo, item)
	}
	return todo
}

func (todos *Todos) GetUpcoming() []Item {
	// return []Item{} // TODO: replace this
	var upcoming []Item
	for _, item := range todos.items {
		if item.Deadline.After(time.Now()) {
			upcoming = append(upcoming, item)
		}
	}
	return upcoming
}

func NewItem(title string, deadline time.Time) Item {
	return Item{title, deadline}
}

func NewTodos() Todos {
	return Todos{}
}

func main() {

	// add some todos
	todos := NewTodos()
	todos.Add(NewItem("Buy milk", time.Date(2023, time.March, 21, 0, 0, 0, 0, time.UTC)))
	todos.Add(NewItem("Buy eggs", time.Date(2022, time.March, 22, 0, 0, 0, 0, time.UTC)))
	todos.Add(NewItem("Buy bread", time.Date(2025, time.March, 23, 0, 0, 0, 0, time.UTC)))

	// print all todos
	allActivities := todos.GetAll()
	i := 1
	fmt.Println("\nTodo's list:")
	for _, item := range allActivities {
		fmt.Printf("%d . %s  %s\n", i, item.Title, item.Deadline)
		i++
	}

	// print upcoming activities
	fmt.Println("\nUpcoming todo's:")
	upcomingActivities := todos.GetUpcoming()
	j := 1
	for _, item := range upcomingActivities {
		fmt.Printf("%d . %s  %s\n", j, item.Title, item.Deadline)
		j++
	}

}
