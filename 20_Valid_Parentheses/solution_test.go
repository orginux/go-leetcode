package validparentheses

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		s := "()"
		want := true
		got := isValid(s)

		if got != want {
			t.Errorf("expected %t but got %t", want, got)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		s := "([)]"
		want := false
		got := isValid(s)

		if got != want {
			t.Errorf("expected %t but got %t", want, got)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		s := "()[]{}"
		want := true
		got := isValid(s)

		if got != want {
			t.Errorf("expected %t but got %t", want, got)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		s := "{[()]}"
		want := true
		got := isValid(s)

		if got != want {
			t.Errorf("expected %t but got %t", want, got)
		}
	})
}
