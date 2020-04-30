package _07_build_tree

import "github.com/mutexlock/algorithms/common"

type TreeNode = common.TreeNode

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := new(TreeNode)
	root.Val = preorder[0]

	index := indexOf(root.Val, inorder)

	root.Left = buildTree(preorder[1:index+1], inorder[0:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])

	return root
}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}

	return 0
}
