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
	"testing"
)

func TestMatrixFloat64Iterator(t *testing.T) {
	columns := 3
	matrix := NewMatrixFloat64(columns)

	err := matrix.AddRow([]float64{1, 2, 3})
	if err != nil {
		t.Errorf("add row error: %+v", err)
	}

	iter := matrix.Iterator()
	for iter.Next() {
		if iter.Index() > 0 {
			t.Errorf("index %d is greater than number of rows", iter.Index())
		}

		row := iter.Row()
		if row.Len() != columns {
			t.Errorf("row length %d does not match number of columns %d", row.Len(), columns)
		}

		indices := iter.RowIndices()
		if len(indices) != columns {
			t.Errorf("row indices length %d does not match number of columns %d", len(indices), columns)
		}

		for _, index := range indices {
			if index < 0 || index > 2 {
				t.Errorf("row index %d is not within bounds", index)
			}
		}
	}

	err = matrix.AddRow([]float64{4, 5, 6})
	if err != nil {
		t.Errorf("add row error: %+v", err)
	}

	iter = matrix.Iterator()
	for iter.Next() {
		if iter.Index() > 1 {
			t.Errorf("index %d is greater than number of rows", iter.Index())
		}

		row := iter.Row()
		if row.Len() != columns {
			t.Errorf("row length %d does not match number of columns %d", row.Len(), columns)
		}

		indices := iter.RowIndices()
		if len(indices) != columns {
			t.Errorf("row indices length %d does not match number of columns %d", len(indices), columns)
		}

		start := iter.Index() * iter.Columns()
		end := start + iter.Columns()
		for _, index := range indices {
			if index < start || index > end {
				t.Errorf("row index %d is not within bound of %d - %d", index, start, end)
			}
		}
	}

	// square
	matrix.Iterator().ApplyToMatrix(func(input interface{}) interface{} {
		v := input.(float64)
		return v * v
	})

	row, err := matrix.GetRow(0)
	if err != nil {
		t.Errorf("get row had error: %+v", err)
	}

	for i := 0; i < row.Len(); i++ {
		if row.Get(i).(float64) != float64((i+1)*(i+1)) {
			t.Errorf("value %v was not square of orignal value %d", row.Get(i).(float64), i+1)
		}
	}

	// column -> add 1
	matrix.Iterator().ApplyToColumns(func(input interface{}) interface{} {
		v := input.(float64)
		return v + 1
	}, []int{0})

	row, err = matrix.GetRow(0)
	if err != nil {
		t.Errorf("get row had error: %+v", err)
	}

	if row.Get(0).(float64) != float64(2) {
		t.Errorf("value %v was not 2", row.Get(0).(float64))
	}
}

func TestMatrixBoolIterator(t *testing.T) {
	columns := 3
	matrix := NewMatrixBool(columns)

	err := matrix.AddRow([]bool{true, true, true})
	if err != nil {
		t.Errorf("add row error: %+v", err)
	}

	iter := matrix.Iterator()
	for iter.Next() {
		if iter.Index() > 0 {
			t.Errorf("index %d is greater than number of rows", iter.Index())
		}

		row := iter.Row()
		if row.Len() != columns {
			t.Errorf("row length %d does not match number of columns %d", row.Len(), columns)
		}

		indices := iter.RowIndices()
		if len(indices) != columns {
			t.Errorf("row indices length %d does not match number of columns %d", len(indices), columns)
		}

		for _, index := range indices {
			if index < 0 || index > 2 {
				t.Errorf("row index %d is not within bounds", index)
			}
		}
	}

	err = matrix.AddRow([]bool{false, false, false})
	if err != nil {
		t.Errorf("add row error: %+v", err)
	}

	iter = matrix.Iterator()
	for iter.Next() {
		if iter.Index() > 1 {
			t.Errorf("index %d is greater than number of rows", iter.Index())
		}

		row := iter.Row()
		if row.Len() != columns {
			t.Errorf("row length %d does not match number of columns %d", row.Len(), columns)
		}

		indices := iter.RowIndices()
		if len(indices) != columns {
			t.Errorf("row indices length %d does not match number of columns %d", len(indices), columns)
		}

		start := iter.Index() * iter.Columns()
		end := start + iter.Columns()
		for _, index := range indices {
			if index < start || index > end {
				t.Errorf("row index %d is not within bound of %d - %d", index, start, end)
			}
		}
	}

	// invert
	matrix.Iterator().ApplyToMatrix(func(input interface{}) interface{} {
		v := input.(bool)

		if v {
			return false
		}

		return true
	})

	row, err := matrix.GetRow(0)
	if err != nil {
		t.Errorf("get row had error: %+v", err)
	}

	for i := 0; i < row.Len(); i++ {
		if row.Get(i).(bool) != false {
			t.Errorf("expected false but value was true")
		}
	}

	// column -> all true
	matrix.Iterator().ApplyToColumns(func(input interface{}) interface{} {
		return true
	}, []int{0})

	row, err = matrix.GetRow(0)
	if err != nil {
		t.Errorf("get row had error: %+v", err)
	}

	if row.Get(0).(bool) != true {
		t.Errorf("value %v was not true", row.Get(0).(bool))
	}
}
