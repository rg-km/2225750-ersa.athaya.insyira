package main

import (
	"errors"
	"fmt"
)

// Dari inisiasi stack dengan jumlah maksimal elemen 10, cobalah implementasikan operasi push.

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
		s.Top += 1
		s.Data = append(s.Data, Elemen)
	}
	return nil
}

func main() {
	// set size stack
	len := 1
	stack := NewStack(len)

	// push some elements
	stack.Push(1)
	err := stack.Push(2)
	if err != nil {
		fmt.Println(err)
	}

	// cetak info stack
	fmt.Println(stack)
	fmt.Println("Top :",stack.Top)
	fmt.Println("Data :",stack.Data)
	fmt.Println("index 0 : ",stack.Data[0])
}
