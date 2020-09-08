package _13_move_count

//地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？
//
//
//
//示例 1：
//
//输入：m = 2, n = 3, k = 1
//输出：3
//示例 2：
//
//输入：m = 3, n = 1, k = 0
//输出：1
//提示：
//
//1 <= n,m <= 100
//0 <= k <= 20

//func movingCount(m int, n int, k int) int {
//	//定义一个m * n 的数组
//	array := make([][]int, m)
//	for i, _ := range array {
//		array[i] = make([]int, n)
//	}
//	//(0,0)必定存在，a从1开始计数
//	array[0][0] = 1
//	a := 1
//	for i := 0; i < m; i++ {
//		for j := 0; j < n; j++ {
//			if (i > 0 && array[i-1][j] == 1) || (j > 0 && array[i][j-1] == 1) {
//				if getres(i, j, k) {
//					array[i][j] = 1
//					a++
//				}
//			}
//		}
//	}
//	return a
//}
//func getres(m, n, k int) bool {
//	res := 0
//	for m > 0 {
//		res += m % 10
//		m /= 10
//	}
//	for n > 0 {
//		res += n % 10
//		n /= 10
//	}
//
//	return res <= k
//}

//func movingCount1(m int, n int, k int) int {
//	// 使用匿名函数 - 需要提前声明才能调用自身
//	var dfs func(y, x int) (res int)
//
//	// 用Map存储已访问过的坐标 - 利用Map的特性
//	visited := make(map[string]bool)
//
//	// 计算数位和
//	digitSum := func(x int) (res int) {
//		for ; x > 0; x = x / 10 {
//			res += x % 10
//		}
//		return
//	}
//
//	// 深度优先搜索
//	dfs = func(y, x int) (res int) {
//		// 判断坐标是否越界
//		if y < 0 || y > m-1 {
//			return
//		}
//		if x < 0 || x > n-1 {
//			return
//		}
//
//		// 判断坐标是否已经访问过
//		if _, isExist := visited[fmt.Sprintf("%d,%d", y, x)]; isExist {
//			return
//		}
//
//		// 判断数位和是否大于k
//		if (digitSum(y) + digitSum(x)) > k {
//			return
//		}
//
//		// 将该坐标加入已访问
//		visited[fmt.Sprintf("%d,%d", y, x)] = true
//		res += 1
//
//		// 搜索相邻的坐标
//		res += dfs(y+1, x)
//		res += dfs(y, x+1)
//		return
//	}
//
//	// 执行DFS算法
//	return dfs(0, 0)
//}

func getSum(m, n int) int {
	sum := 0
	for m != 0 {
		sum += m % 10
		m = m / 10
	}
	for n != 0 {
		sum += n % 10
		n = n / 10
	}
	return sum
}

func dfs(x, y, m, n, k int, vis [][]bool) int {
	if getSum(x, y) > k {
		return 0
	}
	if x >= m || y >= n || x < 0 || y < 0 {
		return 0
	}
	if vis[x][y] {
		return 0
	}
	vis[x][y] = true
	sum := 1
	sum += dfs(x-1, y, m, n, k, vis)
	sum += dfs(x+1, y, m, n, k, vis)
	sum += dfs(x, y-1, m, n, k, vis)
	sum += dfs(x, y+1, m, n, k, vis)
	return sum
}

func movingCount(m int, n int, k int) int {
	vis := make([][]bool, m)
	for i := 0; i < len(vis); i++ {
		vis[i] = make([]bool, n)
	}
	return dfs(0, 0, m, n, k, vis)
}
