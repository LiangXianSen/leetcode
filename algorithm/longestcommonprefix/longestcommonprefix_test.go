package longestcommonprefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	must := assert.New(t)

	cases := []struct {
		Input  []string
		Expect string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"bpad", "bpa", "bp", "bpaf"}, "bp"},
		{[]string{"aaa", "aa", "aaa"}, "aa"},
	}

	for _, v := range cases {
		actual := longestCommonPrefix(v.Input)
		must.Equal(v.Expect, actual)
	}
}

func BenchmarkLongestCommonPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestCommonPrefix([]string{"aaa", "aa", "aaa"})
	}
}
