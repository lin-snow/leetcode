/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

package main

// @lc code=start
func searchInsert(nums []int, target int) int {
    left, right := 0, len(nums)-1
	
	for left <= right {
		mid := left + (right-left)>>1

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid-1
		} else {
			left = mid+1
		}
	}

	return right+1
}
// @lc code=end

