package main

import (
	"fmt"
)

func out(arr [6]int) {
	for i := range arr {
		fmt.Printf("%d ", arr[i])
	}
}

func main() {
	var size = 6                                  // 1
	var arr = [...]int{0, 100, -100, 50, -50, 12} // 2

	for left, right := 0, size-1; left < right; { // 3
		for i := left; i < right; i++ { // 4
			if arr[i+1] < arr[i] { // 5
				temp := arr[i]    // 6
				arr[i] = arr[i+1] // 7
				arr[i+1] = temp   // 8
			}
		}
		right--                         // 9
		for i := right; i > left; i-- { // 10
			if arr[i-1] > arr[i] { // 11
				temp := arr[i]    // 12
				arr[i] = arr[i-1] // 13
				arr[i-1] = temp   // 14
			}
		}
		left++ // 15
	}
	fmt.Printf("Исходный массив: ")
	out(arr)
}

/* 
--- ГУ (CG) ---
1) стрелка 4 -> 3 лишняя (случайно ткнул)
2) из 11 -> 10 не хватает

--- ИГ (IG) ---

*/