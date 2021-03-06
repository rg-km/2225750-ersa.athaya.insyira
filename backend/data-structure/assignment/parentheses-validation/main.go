package main

import (
	"fmt"

	"github.com/ruang-guru/playground/backend/data-structure/assignment/parentheses-validation/stack"
)

// Salah satu problem populer yang dapat diselesaikan dengan menggunakan Stack adalah mengecek validitas tanda kurung.
// Diberikan sebuah string yang hanya terdapat karakter '(', ')', '{', '}', '[', dan ']'.
// Tentukan apakah sebuah string merupakan sekumpulan tanda kurung yang valid.
// String tanda kurung yang valid adalah
// 1. Tanda kurung buka harus ditutup oleh pasangannya.
// 2. Tanda kurung buka harus ditutup di urutan yang benar.

// Contoh 1
// Input: s = "()"
// Output: true
// Penjelasan: tanda kurung buka '(' ditutup dengan pasangannya yaitu ')'.

// Contoh 2
// Input: s = "{)"
// Output: false
// Penjelasan: tanda kurung buka '{' tidak ditutup oleh pasangannya.

// Contoh 3
// Input: s = "([])"
// Output: true
// Penjelasan: tanda kurung buka ditutup dengan pasangannya dan sesuai dengan urutan.

func IsValidParentheses(s string) bool {
	// inisiasi stack
	stack := stack.NewStack()
	
	// jika string kosong, return false
	if s == "" {
		return false
	}
	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack.Push(char)
		} else {
			if stack.IsEmpty() {
				return false
			}
			top, _ := stack.Peek()
			if char == ')' && top != '(' {
				return false
			}
			if char == '}' && top != '{' {
				return false
			}
			if char == ']' && top != '[' {
				return false
			}
			stack.Pop()
		}
	}
	return stack.IsEmpty()
}

func main() {
	fmt.Println(IsValidParentheses("()"))
	fmt.Println(IsValidParentheses("{)"))
	fmt.Println(IsValidParentheses("([])"))
	fmt.Println(IsValidParentheses(""))
}
