package ant

import (
	"fmt"
)

// log - log
func (ant *Ant) log() {
	fmt.Println("Ant pos: ", ant.position)
	for i := range ant.visited {
		fmt.Println()
		for _, k := range ant.visited[i] {
			fmt.Printf("%d ", k)
		}
	}
	fmt.Println("\ncurrent enviroment pheromons")
	for i := range ant.env.pheromon {
		fmt.Println()
		for _, k := range ant.env.pheromon[i] {
			fmt.Printf("%f ", k)
		}
	}
}
