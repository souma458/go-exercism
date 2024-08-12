package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(s string) (Matrix, error) {
	rows := strings.Split(s, "\n")
	matrix := make(Matrix, len(rows))
	expectedColumns := len(strings.Split(strings.TrimSpace(rows[0]), " "))

	for i, row := range rows {
		columns := strings.Split(strings.TrimSpace(row), " ")
		if len(columns) != expectedColumns {
			return Matrix{}, errors.New("invalid input")
		}
		matrix[i] = make([]int, len(columns))
		for j, col := range columns {
			num, err := strconv.ParseInt(col, 10, 64)
			if err != nil {
				return Matrix{}, errors.New("invalid input")
			}
			matrix[i][j] = int(num)
		}
	}
	return matrix, nil
}

func (m Matrix) Cols() [][]int {
	colMatrix := make(Matrix, len(m[0]))
	for _, row := range m {
		for j, col := range row {
			if colMatrix[j] == nil {
				colMatrix[j] = make([]int, 0)
			}
			colMatrix[j] = append(colMatrix[j], col)
		}
	}
	return colMatrix
}

func (m Matrix) Rows() [][]int {
	rowMatrix := make(Matrix, len(m))
	for i, row := range m {
		rowMatrix[i] = make([]int, len(row))
		for j, col := range row {
			rowMatrix[i][j] = int(col)
		}
	}
	return rowMatrix
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 || row >= len(m) || col >= len(m[0]) {
		return false
	}

	m[row][col] = val
	return true
}
