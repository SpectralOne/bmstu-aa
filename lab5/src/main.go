package main

import (
	"fmt"
	"time"

	"./conv"
)

func main() {
	amount := 20

	start := time.Now()
	wait := make(chan int)
	lineP := conv.Parallel(amount, wait)
	<-wait
	// FIXME: uncomment log
	// conv.Log(lineP)
	_ = lineP
	end := time.Now()
	fmt.Printf("%d тортов на параллельном ковейере: %v\n", amount, end.Sub(start))

	start = time.Now()
	lineL := conv.Linear(amount)
	end = time.Now()
	_ = lineL
	// FIXME: uncomment log
	// conv.Log(lineL)
	fmt.Printf("%d тортов на линейном конвейере: %v\n", amount, end.Sub(start))
}
