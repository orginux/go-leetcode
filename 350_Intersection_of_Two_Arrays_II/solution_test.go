package intersectionoftwoarraysii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		nums1 := []int{1, 2, 2, 3}
		nums2 := []int{2, 2}

		want := []int{2, 2}
		got := intersect(nums1, nums2)

		assert.Equal(t, got, want)
	})
	t.Run("Case 2", func(t *testing.T) {
		nums1 := []int{4, 9, 5}
		nums2 := []int{9, 4, 9, 8, 4}

		want := []int{4, 9}
		got := intersect(nums1, nums2)

		assert.Equal(t, got, want)
	})
}
