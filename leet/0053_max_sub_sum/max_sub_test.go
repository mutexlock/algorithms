package _053_max_sub_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_OK(t *testing.T) {
	ast := assert.New(t)
	max := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	ast.Equal(6, max)

}
