package matrix

import "github.com/humilityai/sam"

/*
	The data structure used could have been: map[int64]float64, which would have mapped map[coordinates]value.
	The primary reason that the data structure was not utilized is because the intiial use case for sparse matrix
	needed to allow for simpler column appendage.

	TODO: we can define another column immutable (or dimension immutable) ImmutableSparse Matrix that would utilize the map[coordinates]value structure
*/

// Sparse represents a sparse matrix and should only
// be used when the data in the matrix is expected to be
// mostly zero (unset).
type Sparse struct {
	C    int                     `json:"columns"`
	Data map[int]map[int]float64 `json:"data"`
}

type value struct {
	Column int
	Value  float64
}

type values []value

// NewSparse will return a `*Sparse` matrix.
func NewSparse() *Sparse {
	return &Sparse{
		Data: make(map[int]map[int]float64),
	}
}

// Rows will return the number of rows in the sparse matrix
func (s *Sparse) Rows() int {
	return len(s.Data)
}

// Columns will return the number of columns in the sparse matrix
func (s *Sparse) Columns() int {
	return s.C
}

// Set will set a float64 value at the specified coordinates in
// the matrix.
func (s *Sparse) Set(i, j int, value float64) {
	row, ok := s.Data[i]
	if !ok {
		s.Data[i] = make(map[int]float64)
		row = s.Data[i]
	}

	if j > s.C {
		s.C = j
	}

	row[j] = value
}

// Get will return the value found at the provided coordinates.
// The value will return `0` if the coordinates do not exist
// in the matrix.
func (s *Sparse) Get(i, j int) float64 {
	return s.Data[i][j]
}

// Increment will add +1 to the value found at the coordinates.
// If the coordinates do not exist then they will be created.
func (s *Sparse) Increment(i, j int) {
	row, ok := s.Data[i]
	if !ok {
		s.Data[i] = make(map[int]float64)
		row = s.Data[i]
	}

	if j > s.C {
		s.C = j
	}

	row[j]++
}

// GetRow will return the list of values found at row `i`.
// It will return a list of {column, value} pairs.
func (s *Sparse) GetRow(i int) (values, error) {
	var v values
	row, ok := s.Data[i]
	if !ok {
		return v, ErrRowIndex
	}

	for column, val := range row {
		v = append(v, value{
			Column: column,
			Value:  val,
		})
	}

	return v, nil
}

// Type says the Sparse matrix is a float64 data type.
func (s *Sparse) Type() string {
	return sam.Float64Type
}

// Equal ...
func (v values) Equal(input interface{}) bool { return false }

// Get ...
func (v values) Get(int) interface{} { return nil }

// Len ...
func (v values) Len() int { return len(v) }

// Set ...
func (v values) Set(int, interface{}) {}

// Type ...
func (v values) Type() string { return "" }

// Subslice ...
func (v values) Subslice(start, end int) sam.Slice { return v[start:end] }

// Sum will return the sum of values for a row in a Sparse matrix.
func (v values) Sum() float64 {
	var sum float64
	for _, val := range v {
		sum += val.Value
	}
	return sum
}
