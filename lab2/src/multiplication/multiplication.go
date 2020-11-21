package multiplication

import (
	"../matrix"
	"../types"
)

// Matrix - matrix interface
type Matrix interface {
	types.Matrix
}

// Multiply - standart
func Multiply(a, b Matrix) (res Matrix) {
	rows := a.Rows()
	columns := b.Cols()

	res = matrix.Zeros(rows, columns)

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			for k := 0; k < rows; k++ {
				res.Set(i, j, res.Get(i, j)+a.Get(i, k)*b.Get(k, j))
			}
		}
	}

	return
}

// Winograd - wingrad algorithm
func Winograd(a, b Matrix) (res Matrix) {
	row1 := a.Rows()
	row2 := b.Rows()
	col2 := b.Cols()

	if row2 != a.Cols() {
		return
	}

	rowFactor := make([]int, row1)
	for i := range rowFactor {
		rowFactor[i] = 0
	}
	colFactor := make([]int, col2)
	for i := range rowFactor {
		colFactor[i] = 0
	}

	d := row2 / 2
	for i := 0; i < row1; i++ {
		for j := 0; j < d; j++ {
			rowFactor[i] += a.Get(i, 2*j) * a.Get(i, 2*j+1)
		}
	}

	for i := 0; i < col2; i++ {
		for j := 0; j < d; j++ {
			colFactor[i] += b.Get(2*j, i) * b.Get(2*j+1, i)
		}
	}

	res = matrix.Zeros(row1, col2)
	for i := 0; i < row1; i++ {
		for j := 0; j < col2; j++ {
			res.Set(i, j, -rowFactor[i]-colFactor[j])
			for k := 0; k < d; k++ {
				res.Set(i, j, res.Get(i, j)+
					((a.Get(i, 2*k)+b.Get(2*k+1, j))*
						(a.Get(i, 2*k+1)+b.Get(2*k, j))))
			}
		}
	}

	if row2%2 != 0 {
		for i := 0; i < row1; i++ {
			for j := 0; j < col2; j++ {
				res.Set(i, j, res.Get(i, j)+a.Get(i, row2-1)*b.Get(row2-1, j))
			}
		}
	}

	return
}

// ImpWinograd - improved Winograd
func ImpWinograd(a, b Matrix) (res Matrix) {
	row1 := a.Rows()
	row2 := b.Rows()
	col2 := b.Cols()

	if row2 != a.Cols() {
		return
	}

	rowFactor := make([]int, row1)
	for i := range rowFactor {
		rowFactor[i] = 0
	}
	colFactor := make([]int, col2)
	for i := range rowFactor {
		colFactor[i] = 0
	}

	d := row2 / 2

	for i := 0; i < row1; i++ {
		val := 0
		for j := 0; j < d; j++ {
			val += a.Get(i, 2*j) * a.Get(i, 2*j+1)
		}
		rowFactor[i] = val
	}

	for i := 0; i < col2; i++ {
		val := 0
		for j := 0; j < d; j++ {
			val += b.Get(2*j, i) * b.Get(2*j+1, i)
		}
		colFactor[i] = val
	}

	res = matrix.Zeros(row1, col2)

	for i := 0; i < row1; i++ {
		for j := 0; j < col2; j++ {
			val := 0
			for k := 0; k < d; k++ {
				val += (a.Get(i, 2*k) + b.Get(2*k+1, j)) *
					(a.Get(i, 2*k+1) + b.Get(2*k, j))

			}
			res.Set(i, j, val-rowFactor[i]-colFactor[j])
		}
	}

	if row2%2 != 0 {
		for i := 0; i < row1; i++ {
			for j := 0; j < col2; j++ {
				res.Set(i, j, res.Get(i, j)+a.Get(i, row2-1)*b.Get(row2-1, j))
			}
		}
	}
	return
}
