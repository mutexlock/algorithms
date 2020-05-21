package _07_build_tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	ast.Equal("1", "1", "11")
	tcs := []struct {
		pre, in []int
		ans     [][]int
	}{

		{
			[]int{3, 9, 20, 15, 7},
			[]int{9, 3, 15, 20, 7},
			[][]int{
				[]int{3},
				[]int{9, 20},
				[]int{15, 7},
			},
		},

		// 可以多个 testcase
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		res := Bfs(buildTree(tc.pre, tc.in))
		fmt.Println(res)
		ast.Equal(tc.ans, res, "输入:%v", tc)

		fmt.Println(breadthFirstSearch(buildTree(tc.pre, tc.in)))
	}

	//getIntersectionNode()
}
