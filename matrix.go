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

// Matrix is the minimal interface for
// any new matrices.
type Matrix interface {
	Rows() int
	Columns() int
	GetRow(i int) (sam.Slice, error)
	Type() string
}