package merge_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge_1(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		nums1, m := []int{1, 2, 3, 0, 0, 0}, 3
		nums2, n := []int{2, 5, 6}, 3

		answer := []int{1, 2, 2, 3, 5, 6}
		result := merge(nums1, m, nums2, n)

		assert.Equal(t, result, answer)
	})

	t.Run("Case 2", func(t *testing.T) {

		nums1, m := []int{1}, 1
		nums2, n := []int{}, 0

		answer := []int{1}
		result := merge(nums1, m, nums2, n)

		assert.Equal(t, result, answer)

	})
	t.Run("Case 3", func(t *testing.T) {

		nums1, m := []int{0}, 0
		nums2, n := []int{1}, 1

		answer := []int{1}
		result := merge(nums1, m, nums2, n)

		assert.Equal(t, result, answer)

	})
}
