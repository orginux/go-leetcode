package reshapethematrix

func matrixReshape(mat [][]int, r int, c int) [][]int {
	m := len(mat)    // rows
	n := len(mat[0]) // columns

	// already done
	if m == r && n == n {
		return mat
	}

	// not enough elements
	if m*n < r*c {
		return mat
	}

	// cannot be split
	if m*n%r != 0 {
		return mat
	}

	result := make([][]int, r)
	for i := 0; i < r; i++ {
		result[i] = make([]int, c)
	}

	var resultRow, resultCol int
	for sourceRow := 0; sourceRow < m; sourceRow++ {
		for sourceCol := 0; sourceCol < n; sourceCol++ {
			// fmt.Printf("resultRow: %d, resultCol: %d\n", resultRow, resultCol)
			// fmt.Printf("sourceRow: %d, sourceCol: %d\n", sourceRow, sourceCol)
			// fmt.Println("value: ", mat[sourceRow][sourceCol])
			// fmt.Println()
			result[resultRow][resultCol] = mat[sourceRow][sourceCol]

			// next col
			if resultCol < c-1 {
				resultCol++
			} else { // next row
				resultRow++
				resultCol = 0
			}
		}

	}

	return result
}
