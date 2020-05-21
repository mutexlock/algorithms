package dp

import (
	"fmt"
	"testing"
)

//N = 3, W = 4
//wt = [2, 1, 3]
//val = [4, 2, 3]

func Test_OK(t *testing.T) {
	//	ast := assert.New(t)
	max := MaxValue(4, []int{2, 1, 3}, []int{4, 2, 3}, 3)
	fmt.Println(max)

	fmt.Println(canPartition([]int{2, 1, 3}))
	fmt.Println(canPartition([]int{2, 1, 3, 4}))
	fmt.Println(canPartition([]int{2, 1, 3, 7}))
	//fmt.Println(canPartition([]int{2, 1, 3}))
	//fmt.Println(canPartition([]int{2, 1, 3}))

}
