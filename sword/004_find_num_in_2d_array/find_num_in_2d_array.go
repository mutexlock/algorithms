package _04_find_num_in_2d_array

//在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
//
//
//
//示例:
//
//现有矩阵 matrix 如下：
//
//[
//[1,   4,  7, 11, 15],
//[2,   5,  8, 12, 19],
//[3,   6,  9, 16, 22],
//[10, 13, 14, 17, 24],
//[18, 21, 23, 26, 30]
//]
//给定 target = 5，返回 true。
//
//给定 target = 20，返回 false。
//
//
//
//限制：
//
//0 <= n <= 1000
//
//0 <= m <= 1000

func findNumberIn2DArray(matrix [][]int, target int) bool {
	rows := len(matrix)
	if rows == 0 {
		return false
	}
	columns := len(matrix[0])
	if columns == 0 {
		return false
	}
	//获取左上角左边
	row := 0
	column := columns - 1
	for row < rows && column >= 0 {
		if matrix[row][column] == target {
			return true
		} else if matrix[row][column] > target {
			column--
		} else {
			row++
		}
	}

	return false
}

//class Solution {
//public boolean findNumberIn2DArray(int[][] matrix, int target) {
//if (matrix == null || matrix.length == 0 || matrix[0].length == 0) {
//return false;
//}
//int rows = matrix.length, columns = matrix[0].length;
//int row = 0, column = columns - 1;
//while (row < rows && column >= 0) {
//int num = matrix[row][column];
//if (num == target) {
//return true;
//} else if (num > target) {
//column--;
//} else {
//row++;
//}
//}
//return false;
//}
//}