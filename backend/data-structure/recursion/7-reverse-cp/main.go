// Reverse

package main

import "fmt"

func Reverse(st []string, depth int) string {
	str := ""
	if depth == 0 {
		return st[0]
	} else {
		str = Reverse(st, depth-1)
		return st[depth] + str
	}
}

func main() {
	str := []string{"A", "I", "S", "E", "N", "O", "D", "N", "I"}
	s := Reverse(str, len(str)-1)
	fmt.Println(s)
}
