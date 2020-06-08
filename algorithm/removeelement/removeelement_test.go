package removeelement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	must := assert.New(t)

	cases := []struct {
		Input  []int
		Remove int
		Expect int
	}{
		{[]int{3, 2, 2, 3}, 3, 3},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}

	for _, v := range cases {
		actual := removeElement(v.Input, v.Remove)
		must.Equal(v.Expect, actual)
	}
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeElement([]int{3, 2, 2, 3}, 3)
	}
}
