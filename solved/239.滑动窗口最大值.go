/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 *
 * https://leetcode.cn/problems/sliding-window-maximum/description/
 *
 * algorithms
 * Hard (49.83%)
 * Likes:    3302
 * Dislikes: 0
 * Total Accepted:    1M
 * Total Submissions: 2M
 * Testcase Example:  '[1,3,-1,-3,5,3,6,7]\n3'
 *
 * 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k
 * 个数字。滑动窗口每次只向右移动一位。
 *
 * 返回 滑动窗口中的最大值 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
 * 输出：[3,3,5,5,6,7]
 * 解释：
 * 滑动窗口的位置                最大值
 * ---------------               -----
 * [1  3  -1] -3  5  3  6  7       3
 * ⁠1 [3  -1  -3] 5  3  6  7       3
 * ⁠1  3 [-1  -3  5] 3  6  7       5
 * ⁠1  3  -1 [-3  5  3] 6  7       5
 * ⁠1  3  -1  -3 [5  3  6] 7       6
 * ⁠1  3  -1  -3  5 [3  6  7]      7
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1], k = 1
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 * 1 <= k <= nums.length
 *
 *
 */

package main

type MyQueue struct {
	queue []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		queue: make([]int, 0),
	}
}

func (q *MyQueue) Front() int {
	return q.queue[0]
}

func (q *MyQueue) Back() int {
	return q.queue[len(q.queue)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.queue) == 0
}

func (q *MyQueue) Push(val int) {
	for !q.Empty() && val > q.Back() {
		q.queue = q.queue[:len(q.queue)-1]
	}
	q.queue = append(q.queue, val)
} 

func (q *MyQueue) Pop(val int) {
	if !q.Empty() && val == q.Front() {
		q.queue = q.queue[1:]
	}
}

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	queue := NewMyQueue()
	length := len(nums)

	res := make([]int, 0)
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}

	res = append(res, queue.Front())

	for i := k; i < length; i++ {
		queue.Pop(nums[i-k])

		queue.Push(nums[i])

		res = append(res, queue.Front())
	} 

	return res
}
// @lc code=end

