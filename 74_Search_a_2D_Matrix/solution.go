package main

import "fmt"

func main() {
	matrix := [][]int{[]int{1, 3, 5, 7}, []int{10, 11, 16, 20}, []int{23, 30, 34, 60}}
	target := 3
	fmt.Println(searchMatrix(matrix, target))

}

func searchMatrix(matrix [][]int, target int) bool {
	for _, subArray := range matrix {
		return (binarySearch(subArray, target))
	}
	return false
}

func binarySearch(array []int, target int) {
	for _, n := range subArray {
		if n == target {
			return true
		}
	}
}
