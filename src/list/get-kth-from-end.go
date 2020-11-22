package list

// https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/

// 1->2->3->4->5, k = 2
// when k out range of the max length return original list
func getKthFromEnd(head *ListNode, k int) *ListNode {
	quick := head
	slow := head
	for quick != nil && k > 0 {
		quick = quick.Next
		k--
	}
	for quick != nil {
		quick = quick.Next
		slow = slow.Next
	}
	return slow
}
