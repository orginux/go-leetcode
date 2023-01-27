package containsduplicate

func containsDuplicate(nums []int) bool {
	// max number key-value pairs is len the array
	numsCount := make(map[int]uint8, len(nums))

	for _, num := range nums {
		numsCount[num]++
		if numsCount[num] > 1 {
			return true
		}
	}
	return false
}
