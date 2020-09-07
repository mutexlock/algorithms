package _12_path_metric

//请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。
//
//[["a","b","c","e"],
//["s","f","c","s"],
//["a","d","e","e"]]
//
//但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。
//
//
//
//示例 1：
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
//输出：true
//示例 2：
//
//输入：board = [["a","b"],["c","d"]], word = "abcd"
//输出：false
//提示：
//
//1 <= board.length <= 200
//1 <= board[i].length <= 200

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//如果在数组中找得到第一个数，就执行下一步，否则返回false
			if search(board, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}
func search(board [][]byte, i, j, k int, word string) bool {
	//如果找到最后一个数，则返回true,搜索成功
	if k == len(word) {
		return true
	}
	//i,j的约束条件
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return false
	}

	//进入DFS深度优先搜索
	//先把正在遍历的该值重新赋值，如果在该值的周围都搜索不到目标字符，则再把该值还原
	//如果在数组中找到第一个字符，则进入下一个字符的查找
	if board[i][j] == word[k] {
		temp := board[i][j]
		board[i][j] = ' '
		//下面这个if语句，如果成功进入，说明找到该字符，然后进行下一个字符的搜索,直到所有的搜索都成功，
		//即k == len(word) - 1 的大小时，会返回true，进入该条件语句，然后返回函数true值。
		if search(board, i, j+1, k+1, word) ||
			search(board, i, j-1, k+1, word) ||
			search(board, i+1, j, k+1, word) ||
			search(board, i-1, j, k+1, word) {
			return true
		} else {
			//没有找到目标字符，还原之前重新赋值的board[i][j]
			board[i][j] = temp
		}
	}

	return false
}
