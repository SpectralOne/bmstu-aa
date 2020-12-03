package main

import (
	"sort"
	"fmt"
	"./dict"
)

func main() {
	d := dict.Create(10)

	sort.Slice(d, func(i, j int) bool {
		return d[i]["name"] < d[j]["name"]
	})

	fmt.Printf("\n%v", d)
}
