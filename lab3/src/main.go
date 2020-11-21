package main

import (
	"./sort"
	"fmt"
)

func main() {
	a := sort.Generate(10)

	fmt.Printf("Исходный массив: %v", a)
	sort.BubbleSort(a)
	fmt.Printf("\nСортировка пузырьком: %v", a)
}
