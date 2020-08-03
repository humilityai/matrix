// Copyright 2020 Humility AI Incorporated, All Rights Reserved.
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
	"math"
	"math/rand"

	"github.com/humilityai/sam"
	"gonum.org/v1/gonum/mat"
	"gorgonia.org/tensor"
)

// MatrixFloat64 is backed by a single float64 array.
type MatrixFloat64 struct {
	data    sam.SliceFloat64
	columns int
}

// NewMatrixFloat64 creates a Matrix with the specified column
// count.
func NewMatrixFloat64(columns int) *MatrixFloat64 {
	return &MatrixFloat64{
		data:    make(sam.SliceFloat64, 0),
		columns: columns,
	}
}

// AddRow will append the float64 array to the matrix as a new row.
// If the size of the row does not match the number of columns
// in the matrix then an ErrRowSize will be returned.
func (m *MatrixFloat64) AddRow(row []float64) error {
	if len(row) != m.columns {
		return ErrRowSize
	}

	m.data = append(m.data, row...)

	return nil
}

// Type is the type of values in MatrixFloat64
func (m *MatrixFloat64) Type() string {
	return sam.Float64Type
}

// AppendColumn will add a column to the matrix and place
// the specified default value into each row's column value.
func (m *MatrixFloat64) AppendColumn(defaultValue float64) {
	rows := m.Rows()
	data := make(sam.SliceFloat64, len(m.data)+rows, len(m.data)+rows)

	iter := m.Iterator()
	for iter.Next() {
		indices := iter.RowIndices()
		for _, index := range indices {
			data[index] = m.data[index]
		}
		lastIndex := indices[len(indices)-1]
		data[lastIndex+1] = defaultValue
	}

	m.columns++
	m.data = data
}

// Columns will return the number of columns found
// in the matrix.
func (m *MatrixFloat64) Columns() int {
	return m.columns
}

// Dimensions returns the number of rows and columns
// in the matrix: (rows, columns).
func (m *MatrixFloat64) Dimensions() (int, int) {
	return m.Rows(), m.columns
}

// GetColumnData will return a float64 array that contains all the data points
// of the specified column.
// If the specified column is out of bounds then an ErrColumnIndex will be returned.
func (m *MatrixFloat64) GetColumnData(column int) (data sam.SliceFloat64, err error) {
	if column < 0 || column >= m.columns {
		return data, ErrColumnIndex
	}

	for i := 0; i+m.columns < len(m.data)-1; i += m.columns {
		data = append(data, m.data[i+column])
	}

	return
}

// GetRow ...
func (m *MatrixFloat64) GetRow(row int) (sam.Slice, error) {
	err := m.checkRowAndColumnBounds(row, 0)
	if err != nil {
		return sam.SliceFloat64{}, err
	}
	start := row * m.columns

	return sam.SliceFloat64(m.data[start : start+m.columns]), nil
}

// GetValue will return the float64 value found at the row and column
// arguments provided. It will return an error if something is
// invalid about either the row or column argument.
func (m *MatrixFloat64) GetValue(row, column int) (float64, error) {
	err := m.checkRowAndColumnBounds(row, column)
	if err != nil {
		return 0, err
	}

	return m.data[row*m.columns+column], nil
}

// Iterator will return an object that allows row
// iteration of the matrix.
func (m *MatrixFloat64) Iterator() *Iterator {
	return &Iterator{
		Matrix: m,
		row:    -1,
	}
}

// MaxSum will return the row with the greatest sum
// of its components.
func (m *MatrixFloat64) MaxSum() sam.SliceFloat64 {
	maxSum := math.SmallestNonzeroFloat64
	var max sam.SliceFloat64
	for i := 0; i+m.columns < len(m.data)-1; i += m.columns {
		row := m.data[i : i+m.columns]
		s := sam.SliceFloat64(row)
		if s.Sum() > maxSum {
			max = s
		}
	}

	return max
}

// MinSum will return the row with the smallest sum
// of its components.
func (m *MatrixFloat64) MinSum() sam.SliceFloat64 {
	minSum := math.MaxFloat64
	var min sam.SliceFloat64
	for i := 0; i+m.columns < len(m.data)-1; i += m.columns {
		row := m.data[i : i+m.columns]
		s := sam.SliceFloat64(row)
		if s.Sum() < minSum {
			min = s
		}
	}

	return min
}

// Mode will return the "mode" of each column as
// a "vector" or "row".
func (m *MatrixFloat64) Mode() sam.SliceFloat64 {
	var uniques sam.SliceInt
	uniqueCounts := make(sam.MapIntInt)

	for i := 0; i+m.columns < len(m.data)-1; i += m.columns {
		row := m.data[i : i+m.columns]
		var exists bool
		for _, index := range uniques {
			if sam.SliceFloat64(row).Equal(sam.SliceFloat64(m.data[index : index+m.columns])) {
				uniqueCounts[index]++
				exists = true
			}
		}
		if exists == false {
			uniqueCounts[i] = 1
		}
	}

	k, _ := uniqueCounts.MaxValue()

	return m.data[k : k+m.columns]
}

// NonZeroRows will return a new matrix that contains only the non-zero
// rows of the original matrix.
func (m *MatrixFloat64) NonZeroRows() (*MatrixFloat64, error) {
	matrix := NewMatrixFloat64(m.columns)

	for i := 0; i+m.columns < len(m.data)-1; i += m.columns {
		row := m.data[i : i+m.columns]
		if !sam.SliceFloat64(row).IsZeroed() {
			err := matrix.AddRow(row)
			if err != nil {
				return matrix, err
			}
		}
	}

	return matrix, nil
}

// Rows will return the number of rows found
// in the matrix.
func (m *MatrixFloat64) Rows() int {
	return len(m.data) / m.columns
}

// Sample will grab the number of rows provided as an argument
// randomly. The results are returend as a new *MatrixFloat64.
// If the number of rows is less than zero, then zero rows will
// be returned.
// If the number of rows is equal to or greater than the number of
// rows already in the matrix, then a pointer to the original matrix
// will beb returned.
func (m *MatrixFloat64) Sample(amount int) *MatrixFloat64 {
	sample := NewMatrixFloat64(m.columns)

	if amount < 0 {
		return sample
	} else if amount >= len(m.data) {
		return m
	}

	percentage := (float64(amount) / float64(len(m.data))) * 100

	for i := 0; i+m.columns < len(m.data)-1; i += m.columns {
		row := m.data[i : i+m.columns]
		if float64(rand.Intn(100)) < percentage {
			sample.AddRow(row)
		}
	}

	return sample
}

// SetBackingData will replace the matrix backing array with the
// array provided.
func (m *MatrixFloat64) SetBackingData(data sam.SliceFloat64) {
	m.data = data
}

// ToGonum will create and return a new Gonum Mat64 object
// from the MatrixFloat64
func (m *MatrixFloat64) ToGonum() mat.Matrix {
	return mat.NewDense(m.Rows(), m.Columns(), m.data)
}

// ToTensor will create and return a new Gorgonia Tensor (dense) object
// from the MatrixFloat64.
func (m *MatrixFloat64) ToTensor() tensor.Tensor {
	return tensor.NewDense(tensor.Float64, []int{m.Rows(), m.Columns()}, tensor.WithBacking(m.data))
}

// UpdateValue will update the value found at the provided row and column
// arguments.
// If the row or column are out of bounds for the matrix then the proper
// error will be returned.
func (m *MatrixFloat64) UpdateValue(value float64, row, column int) error {
	err := m.checkRowAndColumnBounds(row, column)
	if err != nil {
		return err
	}

	m.data[row*m.columns+column] = value

	return nil
}

func (m *MatrixFloat64) checkRowAndColumnBounds(row, column int) error {
	rows := len(m.data) / m.columns
	if row > rows || row < 0 {
		return ErrRowIndex
	} else if column < 0 || column > m.columns {
		return ErrColumnIndex
	}

	return nil
}
