package types

// Matrix - matrix interface
type Matrix interface {
	Shape() (rows, cols int)
	
	Rows() (rows int)

	Cols() (cols int)

	Get(row, col int) (elem int)

	Set(row, col, elem int) Matrix 

	Out()
}
