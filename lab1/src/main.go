package main

import (
	"./levenshtein"
	"fmt"
)

func main() {
	fmt.Printf("Введите первое слово: ")
	a := levenshtein.GetLine()
	fmt.Printf("Введите второе слово: ")
	b := levenshtein.GetLine()

	fmt.Println()

	dLevRec := levenshtein.LevRec(a, b)
	fmt.Println("Рекурсивный:")
	fmt.Printf("расстояние: %d\n\n", dLevRec)

	dLevRecMtr, mtrRec := levenshtein.LevMtrRec(a, b)
	fmt.Println("Рекурсивный с заполнением матрицы:")
	levenshtein.OutMatrix(mtrRec)
	fmt.Printf("расстояние: %d\n\n", dLevRecMtr)

	dLevMtr, mtrLev := levenshtein.LevMtr(a, b)
	fmt.Println("Матричный:")
	levenshtein.OutMatrix(mtrLev)
	fmt.Printf("расстояние: %d\n\n", dLevMtr)

	dDamLevMtr, mtrDamLev := levenshtein.DamLevMtr(a, b)
	fmt.Println("Дамерау-Левенштейн:")
	levenshtein.OutMatrix(mtrDamLev)
	fmt.Printf("расстояние: %d\n\n", dDamLevMtr)
}
