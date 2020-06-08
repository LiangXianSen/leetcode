package removedeplicatesfromsortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	must := assert.New(t)

	cases := []struct {
		Input  []int
		Expect int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, v := range cases {
		actual := removeDuplicates(v.Input)
		must.Equal(v.Expect, actual)
	}
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeDuplicates([]int{1, 1, 2})
	}
}
