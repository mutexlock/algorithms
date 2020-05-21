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

//深度DFS

func Dfs(root *TreeNode) {

}

func levelOrderFirst(root *TreeNode, level int) {
	if root == nil {
		return
	}

	// 出现了新的 level
	if level >= len(res) {
		res = append(res, []int{})
	}

	if len(res[level]) == 0 {
		res[level] = append(res[level], root.Val)
	}

	levelOrderFirst(root.Left, level+1)
	levelOrderFirst(root.Right, level+1)
}

func levelOrder(root *TreeNode, level int) {
	if root == nil {
		return
	}

	// 出现了新的 level
	if level >= len(res) {
		res = append(res, []int{})
	}
	res[level] = append(res[level], root.Val)

	levelOrder(root.Left, level+1)
	levelOrder(root.Right, level+1)
}

var res = [][]int{}

func Bfs(root *TreeNode) [][]int {
	levelOrder(root, 0)
	return res
}

//递归
func BfsFirst(root *TreeNode) [][]int {
	levelOrderFirst(root, 0)
	return res
}

//非递归
func breadthFirstSearch(root *TreeNode) []int {
	var result []int
	var nodes []*TreeNode = []*TreeNode{root}

	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		result = append(result, node.Val)
		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
	}
	return result
}
