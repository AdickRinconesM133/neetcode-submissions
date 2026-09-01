/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
	prev := dummy

	curr1 := list1
	curr2 := list2

	for curr1 != nil && curr2 != nil {
		if curr1.Val < curr2.Val {
			prev.Next = curr1
			prev = curr1
			curr1 = curr1.Next
		} else {
			prev.Next = curr2
			prev = curr2
			curr2 = curr2.Next
		}
	}

	if curr1 == nil {
		prev.Next = curr2
	} else {
		prev.Next = curr1
	}
	
	return dummy.Next
}
