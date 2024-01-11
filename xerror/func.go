package main

/*
104. Maximum Depth of Binary Tree

Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf nodes

Example 1:

Input: root = [3,9,20,null,null,15,7]
Output: 3

Example 2:

Input: root = [1,null,2]
Output: 2


Constraints:

The number of nodes in the tree is in the range [0, 104].
-100 <= Node.val <= 100
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left) + 1
	right := maxDepth(root.Right) + 1

	if left > right {
		return left
	}

	return right
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// var dist int

// func amountOfTime(root *TreeNode, start int) int {
// 	dist = 0
// 	depth, infDepth := dfs(root, start)
// 	return max(dist, depth-infDepth)
// }

// func dfs(n *TreeNode, start int) (int, int) {
// 	if n == nil {
// 		return 0, 0
// 	}
// 	d1, infD1 := dfs(n.Left, start)
// 	d2, infD2 := dfs(n.Right, start)
// 	if infD1 != infD2 {
// 		if infD1 > 0 {
// 			dist = max(dist, d2+infD1)
// 		} else {
// 			dist = max(dist, d1+infD2)
// 		}
// 	}
// 	b := max(infD1, infD2)
// 	if n.Val == start || infD1 != infD2 {
// 		b += 1
// 	}
// 	return max(d1, d2) + 1, b
// }

/*
1026. Maximum Difference Between Node and Ancestor

Given the root of a binary tree, find the maximum value v for which there exist different nodes a and b where v = |a.val - b.val| and a is an ancestor of b.

A node a is an ancestor of b if either: any child of a is equal to b or any child of a is an ancestor of b.

Example 1:

Input: root = [8,3,10,1,6,null,14,null,null,4,7,13]
Output: 7
Explanation: We have various ancestor-node differences, some of which are given below :
|8 - 3| = 5
|3 - 7| = 4
|8 - 1| = 7
|10 - 13| = 3
Among all possible differences, the maximum value of 7 is obtained by |8 - 1| = 7.

Example 2:

Input: root = [1,null,2,null,0,3]
Output: 3

Constraints:

The number of nodes in the tree is in the range [2, 5000].
0 <= Node.val <= 105
*/

// func maxAncestorDiff(root *TreeNode) int {

// }
