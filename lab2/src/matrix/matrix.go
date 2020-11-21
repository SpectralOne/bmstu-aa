package matrix

import (
	"fmt"

	"../types"
)

// Matrix - type sctruct
type Matrix struct {
	initialized bool
	shape       *types.Shape
	elements    []int
}

// Out prints matrix to stdout
func (m *Matrix) Out() {
	rows, cols := m.Shape()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%3d ", m.Get(i, j))
		}
		fmt.Printf("\n")
	}
}

// New - ctor
func New(rows, columns int) func(elements ...int) *Matrix {
	constructor := func(elements ...int) *Matrix {
		size := rows * columns
		shape := types.NewShape(rows, columns)

		m := &Matrix{
			initialized: true,
			shape:       shape,
			elements:    make([]int, size),
		}
		copy(m.elements, elements)

		return m
	}

	return constructor
}

// Zeros - return new zeroed matrix R*C size
func Zeros(rows, columns int) *Matrix {
	return New(rows, columns)(make([]int, rows*columns)...)
}

// Shape - return matrix shape
func (m *Matrix) Shape() (rows, columns int) {
	return m.shape.Rows(), m.shape.Cols()
}

// Rows - return rows
func (m *Matrix) Rows() (rows int) {
	rows, _ = m.Shape()
	return rows
}

// Cols - return cols
func (m *Matrix) Cols() (cols int) {
	_, cols = m.Shape()
	return cols
}

// Get - return elem at index
func (m *Matrix) Get(row, column int) (element int) {
	index := row*m.shape.Cols() + column

	return m.elements[index]
}

// Set - updates elem at index
func (m *Matrix) Set(row, column, element int) types.Matrix {
	index := row*m.shape.Cols() + column
	m.elements[index] = element

	return m
}
