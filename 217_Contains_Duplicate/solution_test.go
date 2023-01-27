package containsduplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		nums := []int{1, 2, 3, 1}
		want := true
		got := containsDuplicate(nums)

		if got != want {
			t.Errorf("expected %t but got %t", want, got)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		want := false
		got := containsDuplicate(nums)

		if got != want {
			t.Errorf("expected %t but got %t", want, got)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
		want := true
		got := containsDuplicate(nums)

		if got != want {
			t.Errorf("expected %t but got %t", want, got)
		}
	})
}
