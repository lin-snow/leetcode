/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N 叉树的层序遍历
 *
 * https://leetcode.cn/problems/n-ary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (74.56%)
 * Likes:    515
 * Dislikes: 0
 * Total Accepted:    253.6K
 * Total Submissions: 340.3K
 * Testcase Example:  '[1,null,3,2,4,null,5,6]'
 *
 * 给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。
 *
 * 树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：root = [1,null,3,2,4,null,5,6]
 * 输出：[[1],[3,2,4],[5,6]]
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：root =
 * [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
 * 输出：[[1],[2,3,4,5],[6,7,8,9,10],[11,12,13],[14]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树的高度不会超过 1000
 * 树的节点总数在 [0, 10^4] 之间
 *
 *
 */

package main

import "container/list"

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
    queue := list.New()
    res := [][]int{}       //结果集
    if root == nil{
        return res
    }
    queue.PushBack(root)
    for queue.Len() > 0 {
        length := queue.Len()     //记录当前层的数量
        var tmp []int
        for T := 0; T < length; T++ {        //该层的每个元素：一添加到该层的结果集中；二找到该元素的下层元素加入到队列中，方便下次使用
            myNode := queue.Remove(queue.Front()).(*Node)
            tmp = append(tmp, myNode.Val)
            for i := 0; i < len(myNode.Children); i++ {
                queue.PushBack(myNode.Children[i])
            }
        }
        res = append(res, tmp)
    }

    return res
}
// @lc code=end

