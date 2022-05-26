// Mengecek jika dua string adalah anagram
// Anagram adalah kata yang dibentuk melalui penataan ulang huruf-huruf dari beberapa kata lain.
//
// Contoh 1
// Input: a = "keen", b = "knee"
// Output: "Anagram"
// Explanation: Jika ditata, "knee" dan "keen" memiliki huruf-huruf yang sama, hanya berbeda urutan
//
// Contoh 2
// Input: a = "fried", b = "fired"
// Output: "Anagram"
// Explanation: Jika ditata, "fried" dan "fired" memiliki huruf-huruf yang sama, hanya berbeda urutan
//
// Contoh 3
// Input: a = "apple", b = "paddle"
// Output: "Bukan Anagram"
// Explanation: Jika ditata, "apple" dan "paddle" memiliki huruf-huruf yang berbeda

package main

import "fmt"

func main() {
	var str1 = "apple"
	var str2 = "paddle"
	fmt.Println(AnagramsChecker(str1, str2))
	// AnagramsChecker(str1, str2)

}

func AnagramsChecker(str1 string, str2 string) string {
	var kataInduk []string
	var kataAnak []string
	var kataHasil []string

	// jika beda panjang kata berarti bukan anagram
	if len(str1) != len(str2) {
		return "Bukan Anagram"
	} else {

		// membuat array kata induk
		for _, v := range str1 {
			kataInduk = append(kataInduk, string(v))
		}

		// membuat array kata anak
		for _, v := range str2 {
			kataAnak = append(kataAnak, string(v))
		}

		// membuat array kata hasil dan mencocokan tiap-tiap huruf
		for _, v := range kataInduk {
			for _, v2 := range kataAnak {
				if v == v2 {
					kataHasil = append(kataHasil, v)
					break
				}
			}
		}
	}

	// jika panjang kata hasil sama dengan panjang kata induk maka itu anagram
	if len(kataHasil) == len(kataInduk) {
		return "Anagram"
	}

	return "Bukan Anagram"

	// fmt.Println(kataInduk)
	// fmt.Println(kataAnak)
	// fmt.Println(kataHasil)
}
