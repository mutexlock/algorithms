package _2_common_node_of_link

import "github.com/mutexlock/algorithms/common"

type ListNode = common.ListNode

//输入两个链表，找出它们的第一个公共节点。
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	//求A、B长度
	var lenA, lenB = 1, 1
	h1, h2 := headA, headB
	for h1.Next != nil {
		lenA++
		h1 = h1.Next
	}

	for h2.Next != nil {
		lenB++
		h2 = h2.Next
	}
	diff := 0
	if lenA > lenB {
		diff = lenA - lenB
		for i := 0; i < diff; i++ {
			headA = headA.Next
		}
	} else {
		diff = lenB - lenA
		for i := 0; i < diff; i++ {
			headB = headB.Next
		}
	}

	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}

//
//解题思路：
//我们使用两个指针 node1，node2 分别指向两个链表 headA，headB 的头结点，然后同时分别逐结点遍历，当 node1 到达链表 headA 的末尾时，重新定位到链表 headB 的头结点；当 node2 到达链表 headB 的末尾时，重新定位到链表 headA 的头结点。
//
//这样，当它们相遇时，所指向的结点就是第一个公共结点。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	hasLinkedToB, hasLinkedToA := false, false
	for a != nil && b != nil {
		if a == b {
			return b
		}
		a, b = a.Next, b.Next
		if a == nil && !hasLinkedToB {
			a = headB
			hasLinkedToB = true
		}
		if b == nil && !hasLinkedToA {
			b = headA
			hasLinkedToA = true
		}
	}
	return nil
}
