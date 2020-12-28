package main

import (
	"fmt"
)

func out(arr [4]int) {
	for i := range arr {
		fmt.Printf("%d ", arr[i])
	}
}

func main() {
	var size = 4         // 1
	var arr [4]int       // 2
	left := 0            // 3
	right := size - 1    // 4
	for i := range arr { // 5
		arr[i] = size - i // 6
	}
	for left < right { // 7
		for i := left; i < right; i++ { // 8
			if arr[i+1] < arr[i] { // 9
				temp := arr[i]    // 10
				arr[i] = arr[i+1] // 11
				arr[i+1] = temp   // 12
			}
		}
		right--                         // 13
		for i := right; i > left; i-- { // 14
			if arr[i-1] > arr[i] { // 15
				temp := arr[i]    // 16
				arr[i] = arr[i-1] // 17
				arr[i-1] = temp   // 18
			}
		}
		left++ // 19
	}
	fmt.Printf("Исходный массив: ")
	out(arr)
}
