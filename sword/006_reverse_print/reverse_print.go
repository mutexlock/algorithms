package _06_reverse_print

import "github.com/mutexlock/algorithms/common"

//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
//
//
//
//示例 1：
//
//输入：head = [1,3,2]
//输出：[2,3,1]
//
//
//限制：
//
//0 <= 链表长度 <= 10000

//func reversePrint(head *common.ListNode) []int {
//	return nil
//}
func reversePrint(head *common.ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}
