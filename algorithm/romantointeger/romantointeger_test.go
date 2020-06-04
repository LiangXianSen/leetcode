package romantointeger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	must := assert.New(t)

	cases := []struct {
		Input  string
		Expect int
	}{
		{"III", 3},
		{"IV", 4},
		{"V", 5},
		{"VIII", 8},
		{"IX", 9},
		{"X", 10},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"MCMXCV", 1995},
	}

	for _, v := range cases {
		actual := romanToInt(v.Input)
		must.Equal(v.Expect, actual)
	}
}

func BenchmarkRomanToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		romanToInt("III")
	}
}
