package types

// Shape - dimension (shape) structure type
type Shape struct {
	rows int
	cols int
}

// NewShape - ctor
func NewShape(rows, cols int) (s *Shape) {
	s = &Shape{
		rows: rows,
		cols: cols,
	}

	return
}

// Rows - return rows
func (s *Shape) Rows() int {
	return s.rows
}

// Cols - return cols
func (s *Shape) Cols() int {
	return s.cols
}
