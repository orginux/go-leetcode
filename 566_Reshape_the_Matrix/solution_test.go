package reshapethematrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrixReshape(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {

		mat := [][]int{[]int{1, 2, 3, 4}}
		r := 2
		c := 2
		want := [][]int{[]int{1, 2}, []int{3, 4}}
		got := matrixReshape(mat, r, c)

		assert.Equal(t, got, want)

	})
	t.Run("Case 2", func(t *testing.T) {

		mat := [][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}, []int{13, 14, 15, 16}, []int{17, 18, 19, 20}}
		r := 42
		c := 1
		want := [][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}, []int{13, 14, 15, 16}, []int{17, 18, 19, 20}}
		got := matrixReshape(mat, r, c)

		assert.Equal(t, got, want)

	})
	t.Run("Case 3", func(t *testing.T) {

		mat := [][]int{[]int{1, 2}, []int{3, 4}}
		r := 4
		c := 1
		want := [][]int{[]int{1}, []int{2}, []int{3}, []int{4}}
		got := matrixReshape(mat, r, c)

		assert.Equal(t, got, want)

	})
}
