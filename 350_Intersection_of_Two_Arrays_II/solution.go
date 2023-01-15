package intersectionoftwoarraysii

import (
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	result := []int{}
	var i, j int

	sort.Slice(nums1, func(x, y int) bool {
		return nums1[x] < nums1[y]
	})
	sort.Slice(nums2, func(j, k int) bool {
		return nums2[j] < nums2[k]
	})

	for i < len(nums1) && j < len(nums2) {

		if nums1[i] < nums2[j] {
			i++
			continue
		}
		if nums1[i] > nums2[j] {
			j++
			continue
		}
		result = append(result, nums1[i])
		i++
		j++
	}

	return result
}
