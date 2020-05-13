package addtwonumbers

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sumNode *ListNode
	var nextNode = new(ListNode)
	sumNode = nextNode // 指针引用，即使nextNode被修改，指针指向的位置不变

	var carry int
	for l1 != nil || l2 != nil {
		var v1 int
		var v2 int

		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + carry
		carry = sum / 10
		nextNode.Val = sum % 10

		if l1 == nil && l2 == nil {
			if carry > 0 {
				nextNode.Next = &ListNode{
					Val: carry,
				}
			}
			break
		}
		nextNode.Next = new(ListNode)
		nextNode = nextNode.Next // 修改指针到新的位置
	}
	return sumNode
}
