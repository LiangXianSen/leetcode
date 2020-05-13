package addtwonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	must := assert.New(t)

	cases := []struct {
		NumNode1 *ListNode
		NumNode2 *ListNode
		Expect   *ListNode
	}{
		{
			NumNode1: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
			NumNode2: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			Expect:   &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}},
		},
		{
			NumNode1: &ListNode{Val: 5},
			NumNode2: &ListNode{Val: 5},
			Expect:   &ListNode{Val: 0, Next: &ListNode{Val: 1}},
		},
		{
			NumNode1: &ListNode{Val: 1},
			NumNode2: &ListNode{Val: 9, Next: &ListNode{Val: 9}},
			Expect:   &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}},
		},
	}

	for _, v := range cases {
		actual := addTwoNumbers(v.NumNode1, v.NumNode2)
		must.Equal(v.Expect, actual)
	}
}
