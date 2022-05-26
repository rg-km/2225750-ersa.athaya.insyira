package main

import "fmt"

func main() {
	numList := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	key := int64(9)
	fmt.Println(BinarySearch(numList, key))
}

//Recursive Binary Search
func BinarySearch(numList []int64, key int64) int {
	low := 0
	high := len(numList) - 1

	if low <= high {
		// jika low lebih kecil dari high, maka akan dilakukan proses binary search
		mid := (low + high) / 2
		// jika key lebih kecil dari nilai mid, maka akan dilakukan proses binary search pada low sampai mid
		if key < numList[mid] {
			return BinarySearch(numList[low:mid], key)
		} else if key > numList[mid] {
			// jika key lebih besar dari nilai mid, maka akan dilakukan proses binary search pada mid sampai high
			return BinarySearch(numList[mid+1:high+1], key)
		} else {
			// jika key sama dengan nilai mid, maka akan mengembalikan nilai 1
			return 1
		}
		// TODO: answer here
	}
	return 0
}
