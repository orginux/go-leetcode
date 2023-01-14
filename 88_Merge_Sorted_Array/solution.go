package merge_sorted_array

import (
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) []int {

	// no elements in nums2
	if n == 0 {
		return nums1
	}

	// no elements in nums1
	if m == 0 {
		return nums2
	}

	// merge and sort
	nums1 = append(nums1[:n], nums2...)
	sort.Slice(nums1, func(x, y int) bool {
		return nums1[x] < nums1[y]
	})

	return nums1
}
