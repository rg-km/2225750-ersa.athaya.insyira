// Contoh Array: 23,98,67,45,12,24
// Prosedur: Ambil pasangan pertama dan tukar jika tidak berurutan dan ulangi sampai elemen terakhir diurutkan

// ###### iteration 1:
// step 1: 23,98,67,45,12,24
// step 2: 23,67,98,45,12,24
// step 3: 23,67,45,98,12,24
// step 4: 23,67,45,12,98,24
// step 5: 23,67,45,12,24,98

// result: we got last element sorted!!

// ###### iteration 2:
// step 1: 23,67,45,12,24,98
// step 2: 23,45,67,12,24,98
// step 3: 23,45,12,67,24,98
// step 4: 23,45,12,24,67,98

// result: we got second last element sorted!!

// ###### iteration 3:
// step 1: 23,45,12,24,67,98
// step 2: 23,12,45,24,67,98
// step 3: 23,12,24,45,67,98

// result: we got third last element sorted!!

// ###### iteration 4:
// step 1: 23,12,24,45,67,98
// step 2: 12,23,24,45,67,98

// hasil: kita mendapat array yang diurutkan!!

package main

import "fmt"

func main() {
	var nums = []int{23, 98, 67, 45, 12, 24}
	// var nums = []int{12, 23, 24, 45, 67, 98}
	fmt.Println(sortArray(nums))
}

func sortArray(arr []int) []int {
	fmt.Println("len array : ", len(arr))
	swapped := false //Untuk memeriksa apakah array sudah diurutkan; kemudian return;
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				//elemen bertukar
				// arr[j], arr[j+1] = arr[j+1], arr[j]
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			return arr
		}
	}
	fmt.Println(swapped)
	return arr
}
