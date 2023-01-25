package ransomnote

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanConstruct(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		ransomNote := "a"
		magazine := "b"
		want := false
		got := canConstruct(ransomNote, magazine)
		assert.Equal(t, got, want)
	})
	t.Run("Case 2", func(t *testing.T) {
		ransomNote := "aa"
		magazine := "ab"
		want := false
		got := canConstruct(ransomNote, magazine)
		assert.Equal(t, got, want)
	})
	t.Run("Case 3", func(t *testing.T) {
		ransomNote := "aa"
		magazine := "aab"
		want := true
		got := canConstruct(ransomNote, magazine)
		assert.Equal(t, got, want)
	})
}
