package main

import (
	"fmt"
	"sort"
)

func main() {
	// Cases:

	//nums1, m := []int{1, 2, 3, 0, 0, 0}, 3
	//nums2, n := []int{2, 5, 6}, 3

	// nums1, m := []int{1}, 1
	// nums2, n := []int{}, 0

	nums1, m := []int{0}, 0
	nums2, n := []int{1}, 1

	merge(nums1, m, nums2, n)

}

func merge(nums1 []int, m int, nums2 []int, n int) {

	// no elements in nums2
	if n == 0 {
		fmt.Println(nums1)
		return
	}

	// no elements in nums1
	if m == 0 {
		fmt.Println(nums2[:n])
		return
	}

	// merge and sort
	nums1 = append(nums1[:n], nums2...)
	sort.Slice(nums1, func(x, y int) bool {
		return nums1[x] < nums1[y]
	})

	fmt.Println(nums1)
	return
}
