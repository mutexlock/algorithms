package backtrace

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	ast.Equal("1", "1", "11")

	fmt.Println(solveNQueens(4))

}
