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
type Func func(interface{}) interface{}

// Iterator is an object that can
// be used to traverse the rows of a matrix
// in order exactly once.
type Iterator struct {
	Matrix
	row int
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
func (i *Iterator) Row() sam.Slice {
	row := i.row
	if row < 0 {
		row = 0
	}

	r, _ := i.GetRow(row)
	return r
}

// RowIndices ...
func (i *Iterator) RowIndices() sam.SliceInt {
	row := i.row
	if row < 0 {
		row = 0
	}

	var indices sam.SliceInt
	start := row * i.Columns()
	for j := start; j < start+i.Columns(); j++ {
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
		row := i.Row()
		for i := 0; i < row.Len(); i++ {
			row.Set(i, f(row.Get(i)))
		}
	}
}

// ApplyToColumns can be used to apply a function to the values
// of one or more columns in the matrix.
func (i *Iterator) ApplyToColumns(f Func, columns sam.SliceInt) {
	for i.Next() {
		row := i.Row()
		for i := 0; i < row.Len(); i++ {
			if columns.Contains(i) {
				row.Set(i, f(row.Get(i)))
			}
		}
	}
}
