// Tulis program sebagai fungsi yang memeriksa apakah dua set dikatakan intersection atau tidak.
// Jika kedua set intersection, fungsi tersebut harus mengembalikan nilai intersection nya.
//
// Contoh 1
// Input: a = {"Java", "Python", "Javascript", "C ++", "C#"}, b = {"Swift", "Java", "Kotlin", "Python"}
// Output: {"Java", "Python"}
// Explanation: intersection dari a dan b adalah "Java" dan "Python"
//
// Contoh 2
// Input: a = {"Java", "Python"}, b = {"Swift"}
// Output: {}
// Explanation: tidak ada intersection dari a dan b

package main

import "fmt"

func main() {
	var str1 = []string{"Java", "Python", "Javascript", "C ++", "C#", "Javascript", "Java"}
	var str2 = []string{"Swift", "Java", "Kotlin", "Python", "C#", "C#"}
	fmt.Println("Himpunan A:", str1)
	fmt.Println("Himpunan B:", str2)
	fmt.Println("After Intersection")
	fmt.Println(Intersection(str1, str2))
	fmt.Println("After Remove Duplicate Himpunan B")
	fmt.Println(RemoveDuplicates(str1))
}

func Intersection(str1, str2 []string) (inter []string) {
	// bikin variable penampung
	temp := make(map[string]bool)

	// looping str1[]string untuk mengisi variable penampung
	for _, item := range str1 {
		temp[item] = true
	}

	// looping str2[]string untuk mengecek apakah ada di variable penampung
	// jika ada maka ditambahkan ke variable inter sebagai return
	for _, item := range str2 {
		if _, exist := temp[item]; exist {
			inter = append(inter, item)
		}
	}
	inter = RemoveDuplicates(inter)
	return
}

func RemoveDuplicates(elements []string) (nodups []string) {
	// bikin varible panampung
	temp := make(map[string]bool)

	// looping elements[]string untuk mengisi variable penampung
	// di proses ini duplikat akan dihapus
	for _, item := range elements {
		temp[item] = true
	}

	// looping variable penampung untuk mengambil nilai
	for key := range temp {
		nodups = append(nodups, key)
	}

	return
}
