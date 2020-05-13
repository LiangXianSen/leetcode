package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	must := assert.New(t)

	cases := []struct {
		Input  []int
		Target int
		Expect []int
	}{
		{Input: []int{2, 7, 11, 15}, Target: 9, Expect: []int{0, 1}},
		{Input: []int{2, 5, 5, 11}, Target: 10, Expect: []int{1, 2}},
		{Input: []int{4, 6, 11}, Target: 15, Expect: []int{0, 2}},
	}

	for _, v := range cases {
		actual := twoSum(v.Input, v.Target)
		must.Equal(v.Expect, actual)
	}
}

func TestTwoSumWithLoop(t *testing.T) {
	must := assert.New(t)

	cases := []struct {
		Input  []int
		Target int
		Expect []int
	}{
		{Input: []int{2, 7, 11, 15}, Target: 9, Expect: []int{0, 1}},
		{Input: []int{2, 5, 5, 11}, Target: 10, Expect: []int{1, 2}},
		{Input: []int{4, 6, 11}, Target: 15, Expect: []int{0, 2}},
	}

	for _, v := range cases {
		actual := twoSumWithLoop(v.Input, v.Target)
		must.Equal(v.Expect, actual)
	}
}

// 数据量小看不出优势，数据量小golang中创建map开销大
func BenchmarkTwoSum(b *testing.B) {
	input := []int{2, 7, 11, 15}
	target := 9
	for i := 0; i < b.N; i++ {
		twoSum(input, target)
	}
}

// 无脑暴力循环
func BenchmarkTwoSumWithLoop(b *testing.B) {
	input := []int{2, 7, 11, 15}
	target := 9
	for i := 0; i < b.N; i++ {
		twoSumWithLoop(input, target)
	}
}
