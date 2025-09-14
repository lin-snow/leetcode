/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 *
 * https://leetcode.cn/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (48.21%)
 * Likes:    3377
 * Dislikes: 0
 * Total Accepted:    884.7K
 * Total Submissions: 1.8M
 * Testcase Example:  '"ADOBECODEBANC"\n"ABC"'
 *
 * 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 ""
 * 。
 *
 *
 *
 * 注意：
 *
 *
 * 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
 * 如果 s 中存在这样的子串，我们保证它是唯一的答案。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "ADOBECODEBANC", t = "ABC"
 * 输出："BANC"
 * 解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "a", t = "a"
 * 输出："a"
 * 解释：整个字符串 s 是最小覆盖子串。
 *
 *
 * 示例 3:
 *
 *
 * 输入: s = "a", t = "aa"
 * 输出: ""
 * 解释: t 中两个字符 'a' 均应包含在 s 的子串中，
 * 因此没有符合条件的子字符串，返回空字符串。
 *
 *
 *
 * 提示：
 *
 *
 * ^m == s.length
 * ^n == t.length
 * 1 <= m, n <= 10^5
 * s 和 t 由英文字母组成
 *
 *
 *
 * 进阶：你能设计一个在 o(m+n) 时间内解决此问题的算法吗？
 */

package main

// @lc code=start
func minWindow(s string, t string) string {
	// 默认子串为s
	sLen := len(s)
	l := 0 // 滑动窗口起始位置
	r := sLen-1

	// 统计当前子串的数据
	sMap := make(map[byte]int)
	for i := 0; i < sLen; i++ {
		sMap[s[i]]++
	}

	tMap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		c := t[i]
		tMap[c]++

		if sMap[c] < tMap[c] {
			return ""
		}
	}
	
	// 滑动窗口
	for l <= r {
		lc, rc := false, false

		// 窗口向右缩小
		if tMap[s[l]] == 0 { // 不是
			sMap[s[l]]--
			l++
		} else if tMap[s[l]] < sMap[s[l]] { // 是
			sMap[s[l]]--
			l++
		} else if tMap[s[l]] > sMap[s[l]] {
			return ""
		} else {
			lc = true
		}

		// 窗口向左缩小
		if tMap[s[r]] == 0 { // 不是
			sMap[s[r]]--
			r--
		} else if tMap[s[r]] < sMap[s[r]] { // 是
			sMap[s[r]]--
			r--
		} else if tMap[s[r]] > sMap[s[r]] {
			return ""
		} else {
			rc = true
		}

		if lc && rc {
			break
		}
	}

	// 生成 result
	result := s[l : r+1]

	return result
}
// @lc code=end

