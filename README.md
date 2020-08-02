# `matrix` package

`go get github.com/humilityai/matrix`

Yet another matrix implementation ...

Matrix package provides a simple matrix API for `float64` and `bool` data.

This package was designed to have yet another custom matrix-as-a-data-structure API. It currently does not support matrix operations.

Superior matrix implementations can be found at: https://godoc.org/gorgonia.org/tensor and https://godoc.org/gonum.org/v1/gonum/mat

Or, you can use the API (`ToTensor()` and `ToGonum()`) of this package to easily convert a matrix into a Gorgonia tensor or Gonum matrix.
