package main

import (
	"fmt"

	"github.com/james-bowman/sparse"
)

func main() {
	sparseM := sparse.NewDOK(4, 4)
	sparseM.Set(0, 2, 1)
	sparseM.Set(1, 0, 2)
	sparseM.Set(2, 3, 3)
	sparseM.Set(3, 1, 4)

	fmt.Print("DOK Matrix: \n", sparseM, "\n\n")
	fmt.Print("CSR Matrix:\n", sparseM.ToCSR(), "\n\n")
	fmt.Print("CSC Matrix:\n", sparseM.ToCSC(), "\n\n")
}
