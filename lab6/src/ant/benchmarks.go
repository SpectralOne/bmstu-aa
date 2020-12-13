package ant

import (
	"fmt"
	"os"
	"time"
)

// BenchAnt - benchmark
func BenchAnt() {
	ants := make([]time.Duration, 0)
	for i := 1; i < 20; i++ {
		_ = os.Remove("data.txt")
		GenData(10)
		env := newEnv("data.txt")
		env.p = 1.0 / float64(i)
		start := time.Now()
		shortest := StartAnt(env, 100)
		end := time.Now()
		ants = append(ants, end.Sub(start))

		_ = shortest
	}

	fmt.Println("Ants algotithm time")
	for j := 0; j < len(ants); j++ {
		fmt.Println(j, ants[j])
	}
}

// BenchAll - benchmark
func BenchAll() {
	ants := make([]time.Duration, 0)
	Brutes := make([]time.Duration, 0)
	for i := 2; i < 11; i++ {
		_ = os.Remove("data.txt")
		GenData(i)
		env := newEnv("data.txt")
		start := time.Now()
		shortest := StartAnt(env, 100) //fmt.Print(shortest, " ")
		end := time.Now()
		ants = append(ants, end.Sub(start))

		start = time.Now()
		shortest = Brute("data.txt") //fmt.Print(shortest, "\n")
		end = time.Now()
		Brutes = append(Brutes, end.Sub(start))
		_ = shortest
	}

	fmt.Println("Ants algotithm time")
	for j := 0; j < len(ants); j++ {
		fmt.Println(j+2, ants[j])
	}
	fmt.Println("Brute force time")
	for j := 0; j < len(ants); j++ {
		fmt.Println(j+2, Brutes[j])
	}
}