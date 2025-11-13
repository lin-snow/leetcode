/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 *
 * https://leetcode.cn/problems/spiral-matrix-ii/description/
 *
 * algorithms
 * Medium (70.31%)
 * Likes:    1490
 * Dislikes: 0
 * Total Accepted:    573.6K
 * Total Submissions: 815.9K
 * Testcase Example:  '3'
 *
 * 给你一个正整数 n ，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：[[1,2,3],[8,9,4],[7,6,5]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：[[1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 *
 *
 */

package main

// @lc code=start
func generateMatrix(n int) [][]int {
	startx, starty := 0, 0
	var loop int = n / 2 // n为偶数时候的循环的次数
	var center int = n / 2
	count := 1
	offset := 1
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	for loop > 0 {
		i, j := startx, starty

		for j = starty; j < n-offset; j++ {
			res[startx][j] = count
			count++
		}

		for i = startx; i < n-offset; i++ {
			res[i][j] = count
			count++
		}

		for ; j > starty; j-- {
			res[i][j] = count
			count++
		}

		for ; i > startx; i-- {
			res[i][j] = count
			count++
		}

		startx++
		starty++
		offset++
		loop--
	}
	if n % 2 == 1 {
		res[center][center] = n * n
	}

	return res
}
// @lc code=end

