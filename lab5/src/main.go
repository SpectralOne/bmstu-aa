package main

import (
	"fmt"
	"time"

	"./conv"
)

func main() {
	amount := 20
	wait := make(chan int)

	start := time.Now()
	lineP := conv.Parallel(amount, wait)
	<-wait
	// FIXME: uncomment log
	conv.Log(lineP, false)
	_ = lineP
	end := time.Now()
	fmt.Printf("%d персональных карточек на пареллельном конвейере: %v\n", amount, end.Sub(start))

	start = time.Now()
	lineL := conv.Linear(amount)
	end = time.Now()
	_ = lineL
	// FIXME: uncomment log
	conv.Log(lineL, false)
	fmt.Printf("%d персональных карточек на линейном конвейере: %v\n", amount, end.Sub(start))
}
