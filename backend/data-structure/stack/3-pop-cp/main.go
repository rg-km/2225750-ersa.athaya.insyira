package main

import (
	"errors"
	"fmt"
)

// Dari inisiasi stack dengan maksimal elemen sebanyak 10, implementasikan operasi pop.

var ErrStackUnderflow = errors.New("stack underflow")
var ErrStackOverflow = errors.New("stack overflow")

type Stack struct {
	Size int
	Top  int
	Data []int
}

func NewStack(size int) Stack {
	return Stack{
		Size: size,
		Top:  -1,
		Data: make([]int, 0),
	}
}

func (s *Stack) Push(Elemen int) error {
	// cek apakah masih ada tempat untuk menambah elemen, jika ga ada maka error stackoverflow
	if s.Top >= s.Size-1 {
		return ErrStackOverflow
	} else {
		s.Top++
		s.Data = append(s.Data, Elemen)
		return nil
	}
}

func (s *Stack) Pop() (int, error) {
	// cek apakah stack kosong, jika kosong maka error stackunderflow
	if s.IsEmpty() {
		return 0, ErrStackUnderflow
	} else {
		// ambil elemen dari stack
		e := s.Data[s.Top]
		s.Top--
		s.Data = s.Data[:s.Top+1]
		return e, nil
	}
}

// function untuk cek apakah stack kosong
func (s *Stack) IsEmpty() bool {
	return s.Top == -1
}

func main() {
	// set size stack
	len := 0
	stack := NewStack(len)

	// push some elements
	for i := 0; i < len; i++ {
		err := stack.Push(i)
		if err != nil {
			fmt.Println(err)
		}
	}

	// hapus elemen teratas
	_, err := stack.Pop()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(stack)
}
