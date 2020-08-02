// Copyright 2020 Hummility AI Incorporated, All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package matrix

import (
	"github.com/humilityai/sam"
	"gorgonia.org/tensor"
)

// MatrixBool is backed by a single array.
type MatrixBool struct {
	data    sam.SliceBool
	columns int
}

// NewMatrixBool creates a Matrix with the specified column
// count.
func NewMatrixBool(columns int) *MatrixBool {
	return &MatrixBool{
		data:    make(sam.SliceBool, 0),
		columns: columns,
	}
}

// AddRow will append the boolean array to the matrix as a new row.
// If the size of the row does not match the number of columns
// in the matrix then an ErrRowSize will be returned.
func (m *MatrixBool) AddRow(row sam.SliceBool) error {
	if len(row) != m.columns {
		return ErrRowSize
	}

	m.data = append(m.data, row...)

	return nil
}

// Columns will return the number of columns found
// in the matrix.
func (m *MatrixBool) Columns() int {
	return m.columns
}

// Dimensions returns the number of rows and columns
// in the matrix: (rows, columns).
func (m *MatrixBool) Dimensions() (int, int) {
	return len(m.data), m.columns
}

// GetColumnData will return a float64 array that contains all the data points
// of the specified column.
// If the specified column is out of bounds then an ErrColumnIndex will be returned.
func (m *MatrixBool) GetColumnData(column int) (data sam.SliceBool, err error) {
	if column < 0 || column >= m.columns {
		return data, ErrColumnIndex
	}

	for i := 0; i+m.columns < len(m.data)-1; i += m.columns {
		data = append(data, m.data[i+column])
	}

	return
}

// GetRow ...
func (m *MatrixBool) GetRow(row int) (sam.SliceBool, error) {
	err := m.checkRowAndColumnBounds(row, 0)
	if err != nil {
		return []bool{}, err
	}
	start := row * m.columns

	return m.data[start : start+m.columns], nil
}

// GetValue will return the boolean value found at the row and column
// arguments provided. It will return an error if something is
// invalid about either the row or column argument.
func (m *MatrixBool) GetValue(row, column int) (bool, error) {
	err := m.checkRowAndColumnBounds(row, column)
	if err != nil {
		return false, err
	}

	return m.data[row*m.columns+column], nil
}

// Rows will return the number of rows found
// in the matrix.
func (m *MatrixBool) Rows() int {
	return len(m.data) / m.columns
}

// SetBackingData will replace the matrix backing array with the
// array provided.
func (m *MatrixBool) SetBackingData(data sam.SliceBool) {
	m.data = data
}

// ToTensor will create and return a new Gorgonia Tensor (dense) object
// from the MatrixBool.
func (m *MatrixBool) ToTensor() tensor.Tensor {
	return tensor.NewDense(tensor.Bool, []int{m.Rows(), m.Columns()}, tensor.WithBacking(m.data))
}

// Type is the type of values in MatrixFloat64
func (m *MatrixBool) Type() string {
	return sam.BoolType
}

// UpdateValue will update the value found at the provided row and column
// arguments.
// If the row or column are out of bounds for the matrix then the proper
// error will be returned.
func (m *MatrixBool) UpdateValue(value bool, row, column int) error {
	err := m.checkRowAndColumnBounds(row, column)
	if err != nil {
		return err
	}

	m.data[row*m.columns+column] = value

	return nil
}

func (m *MatrixBool) checkRowAndColumnBounds(row, column int) error {
	rows := len(m.data) / m.columns
	if row > rows || row < 0 {
		return ErrRowIndex
	} else if column < 0 || column > m.columns {
		return ErrColumnIndex
	}

	return nil
}
