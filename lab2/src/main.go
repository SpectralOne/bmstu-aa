package main

import (
	"./matrix"
	"./multiplication"
	
	"fmt"
)

func main() {
	a := matrix.New(3, 3) (
		1, 2, 3,
		4, 5, 6,
		7, 8, 9)
	b := matrix.New(3, 3)(
		1, 0, 0,
		0, 1, 0,
		0, 0, 1)
	
	fmt.Println("Матрица А:")
	a.Out()
	fmt.Println("Матрица Б:")
	b.Out()

	c := multiplication.Multiply(a, b)
	fmt.Println("Классическое умножение:")
	c.Out()

	d := multiplication.Winograd(a, b)
	fmt.Println("Алгоритм Винограда:")
	d.Out()

	e := multiplication.ImpWinograd(a, b)
	fmt.Println("Оптимизированный алгоритм Винограда:")
	e.Out()
}
