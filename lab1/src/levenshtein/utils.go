package levenshtein

import "fmt"

// OutMatrix prints matrix to stdout
func OutMatrix(mtr [][]int) {
	fmt.Println("Result matrix:")
	for i := range mtr {
		for j := range mtr[i] {
			fmt.Printf("%3d ", mtr[i][j])
		}
		fmt.Printf("\n")
	}
}

// createMatrix is used to create and fill matrix for levenstein
func createMatrix(n, m int, fill bool) (mtr [][]int) {
	mtr = make([][]int, n)
	for i := range mtr {
		mtr[i] = make([]int, m)
		mtr[i][0] = i
	}

	for j := 1; j < m; j++ {
		mtr[0][j] = j
	}

	if fill {
		for i := range mtr {
			for j := range mtr[i] {
				mtr[i][j] = -1
			}
		}
	}

	return
}

// minOf finds minimum from sequence of ints
func minOf(args ...int) (min int) {
	min = args[0]

	for _, val := range args {
		if min > val {
			min = val
		}
	}

	return
}

// GetLine reads line from stdin
func GetLine() string {
	var input string
	fmt.Scanln(&input)

	return input
}
