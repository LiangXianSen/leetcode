package reverseinteger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	must := assert.New(t)

	cases := []struct {
		Input  int
		Expect int
	}{
		{123, 321},
		{456, 654},
		{1 << 31, 0},
		{-1 << 31, 0},
	}

	for _, v := range cases {
		actual := reverse(v.Input)
		must.Equal(v.Expect, actual)
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse(123)
	}
}
