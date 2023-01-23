package firstuniquecharacterinastring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstUniqChar(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		s := "leetcode"
		want := 0
		got := firstUniqChar(s)
		assert.Equal(t, got, want)
	})
	t.Run("Case 2", func(t *testing.T) {
		s := "loveleetcode"
		want := 1
		got := firstUniqChar(s)
		assert.Equal(t, got, want)
	})
	t.Run("Case 3", func(t *testing.T) {
		s := "aabb"
		want := -1
		got := firstUniqChar(s)
		assert.Equal(t, got, want)
	})
}
