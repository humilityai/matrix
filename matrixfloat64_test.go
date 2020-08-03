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

	"github.com/humilityai/sam"
)

func TestMatrixFloat64(t *testing.T) {
	columns := 3
	matrix := NewMatrixFloat64(columns)

	if matrix.Columns() != columns {
		t.Errorf("matrix columns %d does not match columns argument %d", matrix.Columns(), columns)
	}

	if matrix.Rows() != 0 {
		t.Errorf("rows is not 0")
	}

	err := matrix.AddRow([]float64{1, 2, 3})
	if err != nil {
		t.Errorf("matrix row add error: %+v", err)
	}

	if matrix.Rows() != 1 {
		t.Errorf("rows is not 1")
	}

	err = matrix.AddRow([]float64{1, 2})
	if err != ErrRowSize {
		t.Errorf("matrix ErrRowSize was not caught")
	}

	if matrix.Type() != sam.Float64Type {
		t.Errorf("matrix type %v is not %v", matrix.Type(), sam.Float64Type)
	}

	r, c := matrix.Dimensions()
	if r != 1 {
		t.Errorf("rows is %d and not %d", r, matrix.Rows())
	}

	if c != 3 {
		t.Errorf("columns is %d and not %d", c, columns)
	}

	matrix.AppendColumn(0)

	if matrix.Columns() != columns+1 {
		t.Errorf("columns is %d and not %d", matrix.Columns(), columns+1)
	}
}
