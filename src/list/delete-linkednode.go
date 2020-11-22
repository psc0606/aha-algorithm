package list

// https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/
// delete a node from list
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		head = head.Next
		return head
	}
	p := head.Next
	q := head
	for p != nil {
		if p.Val == val {
			q.Next = p.Next
			break
		}
		q = p
		p = p.Next
	}
	return head
}
