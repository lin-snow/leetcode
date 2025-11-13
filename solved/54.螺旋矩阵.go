/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 *
 * https://leetcode.cn/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (54.15%)
 * Likes:    1989
 * Dislikes: 0
 * Total Accepted:    823.5K
 * Total Submissions: 1.5M
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[1,2,3,6,9,8,7,4,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
 * 输出：[1,2,3,4,8,12,11,10,9,5,6,7]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1
 * -100
 *
 *
 */

package main

// @lc code=start
func spiralOrder(matrix [][]int) []int {
    startx, starty := 0, 0
	var m, n int
	if len(matrix) == 0 {
		return []int{}
	} else {
		m = len(matrix)
	}
	if matrix[0] != nil {
		n = len(matrix[0])
	}
	total := m * n
	res := make([]int, total)
	count := total
	offset := 1

	for count > 0 {
		i, j := startx, starty

		// 仅剩一行
		if startx == m-offset && starty <= n-offset {
			for j = starty; j <= n-offset && count > 0; j++ {
				res[total-count] = matrix[i][j]
				count--
			}
			break
		}
		// 仅剩一列
		if starty == n-offset && startx <= m-offset {
			for i = startx; i <= m-offset && count > 0; i++ {
				res[total-count] = matrix[i][j]
				count--
			}
			break
		}

		// 上边：左 -> 右（不含最右角）
		for j = starty; j < n-offset && count > 0; j++ {
			res[total-count] = matrix[i][j]
			count--
		}
		// 右边：上 -> 下（不含最下角）
		for i = startx; i < m-offset && count > 0; i++ {
			res[total-count] = matrix[i][j]
			count--
		}
		// 下边：右 -> 左（不含最左角）
		for ; j > starty && count > 0; j-- {
			res[total-count] = matrix[i][j]
			count--
		}
		// 左边：下 -> 上（不含最上角）
		for ; i > startx && count > 0; i-- {
			res[total-count] = matrix[i][j]
			count--
		}

		offset++
		startx++
		starty++
	}

	return res
}
// @lc code=end

