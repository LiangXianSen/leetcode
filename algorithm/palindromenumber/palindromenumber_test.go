package palindromenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindromNumber(t *testing.T) {
	must := assert.New(t)

	cases := []struct {
		Input  int
		Expect bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{11, true},
		{12321, true},
	}

	for _, v := range cases {
		actual := isPalindrome(v.Input)
		must.Equal(v.Expect, actual)
	}
}

func TestSplitNum(t *testing.T) {
	must := assert.New(t)

	cases := []struct {
		Input  int
		Expect []int
	}{
		{121, []int{1, 2, 1}},
		{-121, []int{-1, -2, -1}},
		{10, []int{0, 1}},
		{11, []int{1, 1}},
		{12321, []int{1, 2, 3, 2, 1}},
	}

	for _, v := range cases {
		actual := splitNum(v.Input)
		must.Equal(v.Expect, actual)
	}
}

func BenchmarkPalindromNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome(121)
	}
}
