package mergetwosortedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	must := assert.New(t)

	cases := []struct {
		NumNode1 *ListNode
		NumNode2 *ListNode
		Expect   *ListNode
	}{
		{
			NumNode1: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
			NumNode2: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
			Expect:   &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4}}}}}},
		},
	}

	for _, v := range cases {
		actual := mergeTwoLists(v.NumNode1, v.NumNode2)
		must.Equal(v.Expect, actual)
	}
}
