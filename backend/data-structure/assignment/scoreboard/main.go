package main

import (
	"fmt"
	"sort"
)

type Score struct {
	Name    string
	Correct int
	Wrong   int
	Empty   int
}
type Scores []Score

func (s Scores) Len() int {
	return len(s)
}

func (s Scores) Less(i, j int) bool {

	// hitung score
	score1 := s[i].Correct*4 - s[i].Wrong
	score2 := s[j].Correct*4 - s[j].Wrong

	// kalau score1 lebih besar, return true 
	// kalau score1 lebih kecil, return false
	if score1 > score2 {
		return true
	} else if score1 < score2 {
		return false
	}

	// Yang Jumlah Benar-nya lebih tinggi akan diurutkan di atas.
	if s[i].Correct > s[j].Correct {
		return true
	} else if s[i].Correct < s[j].Correct {
		return false
	}

	return false
}

func (s Scores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Scores) TopStudents() []string {
	sort.Sort(s) // sorting
	names := []string{}
	for _, score := range s {
		names = append(names, score.Name)
	}
	return names
}

func main() {
	scores := Scores([]Score{
		{"Levi", 3, 2, 2},
		{"Agus", 5, 2, 0},
		{"Doni", 3, 0, 7},
		{"Ega", 3, 0, 7},
		{"Anton", 2, 0, 5},
	})
	sort.Sort(scores)
	fmt.Println(scores.TopStudents())
}
