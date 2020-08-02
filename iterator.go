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
)

// Func is the basic function type for
// modifying the values of a matrix.
type Func func(float64) float64

// Iterator is an object that can
// be used to traverse the rows of a matrix
// in order exactly once.
type Iterator struct {
	*MatrixFloat64
	row int
}

// Iterator will return an object that allows row
// iteration of the matrix.
func (m *MatrixFloat64) Iterator() *Iterator {
	return &Iterator{
		MatrixFloat64: m,
		row:           -1,
	}
}

// Next will set the iterator to return the next row.
// It returns false if the row is larger than the number
// of rows in the matrix.
func (i *Iterator) Next() bool {
	i.row++

	if i.row > i.Rows() {
		return false
	}

	return true
}

// Row will return the data of the current row for the iterator.
func (i *Iterator) Row() sam.SliceFloat64 {
	row := i.row
	if row < 0 {
		row = 0
	}

	start := row * i.columns
	return i.data[start : start+i.columns]
}

// RowIndices ...
func (i *Iterator) RowIndices() sam.SliceInt {
	row := i.row
	if row < 0 {
		row = 0
	}

	var indices sam.SliceInt
	start := row * i.columns
	for j := start; j < start+i.columns; j++ {
		indices = append(indices, j)
	}

	return indices
}

// ApplyToMatrix will apply the supplied function
// to all values in the matrix.
// This should only be called after the Iterator has been created
// and before the Next() method has been called
func (i *Iterator) ApplyToMatrix(f Func) {
	for i.Next() {
		row := i.data[i.row*i.columns : i.row*i.columns+i.columns]
		for i, v := range row {
			row[i] = f(v)
		}
	}
}

// ApplyToColumns can be used to apply a function to the values
// of one or more columns in the matrix.
func (i *Iterator) ApplyToColumns(f Func, columns []int) {
	for i.Next() {
		row := i.data[i.row*i.columns : i.row*i.columns+i.columns]
		for i, v := range row {
			for _, column := range columns {
				if i == column {
					row[i] = f(v)
				}
			}
		}
	}
}
