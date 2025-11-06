/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 找出字符串中第一个匹配项的下标
 *
 * https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/description/
 *
 * algorithms
 * Easy (45.06%)
 * Likes:    2484
 * Dislikes: 0
 * Total Accepted:    1.4M
 * Total Submissions: 3M
 * Testcase Example:  '"sadbutsad"\n"sad"'
 *
 * 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0
 * 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：haystack = "sadbutsad", needle = "sad"
 * 输出：0
 * 解释："sad" 在下标 0 和 6 处匹配。
 * 第一个匹配项的下标是 0 ，所以返回 0 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：haystack = "leetcode", needle = "leeto"
 * 输出：-1
 * 解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= haystack.length, needle.length <= 10^4
 * haystack 和 needle 仅由小写英文字符组成
 *
 *
 */

package main

func getNext(next []int, s string) {
	j := 0
	next[0] = j

	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j - 1]
		}

		if s[i] == s[j] {
			j++
		}

		next[i] = j
	}
}

// @lc code=start
func strStr(haystack string, needle string) int {
    n := len(needle)
	if n == 0 {
		return 0
	}

	j := 0
	next := make([]int, n)
	getNext(next, needle)

	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}

		if haystack[i] == needle[j] {
			j++
		}

		if j == n {
			return i - j + 1
		}
	}
	
	return -1
}
// @lc code=end

