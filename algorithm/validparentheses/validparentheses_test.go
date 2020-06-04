package validparentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	must := assert.New(t)

	cases := []struct {
		Input  string
		Expect bool
	}{
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}

	for _, v := range cases {
		actual := isValid(v.Input)
		must.Equal(v.Expect, actual)
	}
}

func BenchmarkIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isValid("()[]{}")
	}
}
