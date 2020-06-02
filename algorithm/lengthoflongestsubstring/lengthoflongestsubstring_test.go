package lengthoflongestsubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	must := assert.New(t)

	cases := []struct {
		Input  string
		Expect int
	}{
		{"abcabcbb", 3},
		{"bbbbbb", 1},
		{"pwwkeww", 3},
		{"abba", 2},
		{"dvdf", 3},
		{"pwwkewwdsfgc", 6},
		{"aab", 2},
		{"cdd", 2},
		{"abcdefg", 7},
		{"au", 2},
		{" ", 1},
		{"你好 中国", 5},
	}

	for _, v := range cases {
		actual := lengthOfLongestSubstring(v.Input)
		must.Equal(v.Expect, actual)
	}
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring("abcabcbb")
	}
}

func BenchmarkLengthOfLongestSubstringChinese(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring("你好 中国")
	}
}
