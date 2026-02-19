package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func prettyPrintMatrix(m mat.Matrix) {
	formattedM := mat.Formatted(m, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", formattedM)
}

func main() {
	m2d := mat.NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	m2d.Add(m2d, m2d)
	prettyPrintMatrix(m2d)
}
