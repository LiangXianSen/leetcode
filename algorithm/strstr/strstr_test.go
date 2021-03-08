package strstr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	must := assert.New(t)

	cases := []struct {
		Input  []string
		Expect int
	}{
		{[]string{"hello", "ll"}, 2},
		{[]string{"aaa", "a"}, 0},
		{[]string{"aaa", "aaaa"}, -1},
		{[]string{"mississippi", "issip"}, 4},
	}

	for _, v := range cases {
		actual := strStr(v.Input[0], v.Input[1])
		must.Equal(v.Expect, actual)
	}
}

func BenchmarkStrStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strStr("hello", "ll")
	}
}
