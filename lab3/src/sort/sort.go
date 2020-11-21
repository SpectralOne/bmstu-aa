package sort

import (
	"math/rand"
)

// BubbleSort - bubble sort
func BubbleSort(a []int) {
	for range a {
		for i := len(a) - 1; i > 0; i-- {
			if a[i-1] > a[i] {
				a[i-1], a[i] = a[i], a[i-1]
			}
		}
	}
}

// InsertSort - insert sort
func InsertSort(a []int) {
	for i := 1; i < len(a); i++ {
		key := a[i]
		j := i - 1
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = key
	}
}

// Unused due to Stroganov's task (write normal qs)
func partition(a []int, low, high int) int {
	i := low - 1
	pivot := a[high]
	for j := low; j < high; j++ {
		if a[j] <= pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
		a[i+1], a[high] = a[high], a[i+1]
	}
	return i + 1
}

// QuickSort - quick sort
func QuickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivotIndex := rand.Int() % len(a)

	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	QuickSort(a[:left])
	QuickSort(a[left+1:])

	return a
}

var set = [...]int{0, 100, -100, 50, -50}

// Generate - Stroganov case
func Generate(size int) (a []int) {
	a = make([]int, size)
	for i := range a {
		a[i] = set[rand.Intn(4)]
	}

	return
}

// CreateRandom - create
func CreateRandom(size int) (a []int) {
	a = make([]int, size)
	for i := range a {
		a[i] = rand.Intn(100)
	}

	return
}

// CreateSorted - create
func CreateSorted(size int, reverse bool) (a []int) {
	a = make([]int, size)

	for i := range a {
		if reverse {
			a[i] = size - i
		} else {
			a[i] = i
		}
	}

	return
}
