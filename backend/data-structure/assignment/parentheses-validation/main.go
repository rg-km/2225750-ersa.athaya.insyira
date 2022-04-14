package main

import (
	"strings"

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
	if s == "" {
		return false
	}
	stack := new(stack.Stack)
	stack.Top = -1

	tandaKurung := map[string]int{
		"{": 1, "}": -1, "(": 2, ")": -2, "[": 3, "]": -3,
	}

	for i := 0; i < strings.Count(s, "")-1; i++ {
		if tandaKurung[string(s[i])] > 0 {
		} else {
			peekValue, err := stack.Peek()
			if err == nil {
				if tandaKurung[string(s[i])]+tandaKurung[string(peekValue)] == 0 {
					stack.Pop()
				}
			}
		}
	}
	if len(stack.Data) == 0 && tandaKurung[string(s[0])] > 0 {
		return true
	}
	return false // TODO: answer here
}
