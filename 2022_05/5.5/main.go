package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	maxSum int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	partMaxSum := root.Val
	left := findMaxSum(root.Left)
	if left > 0 {
		partMaxSum += left
	}
	right := findMaxSum(root.Right)
	if right > 0 {
		partMaxSum += right
	}
	maxSum = max(maxSum, partMaxSum)
	return max(root.Val, max(root.Val+left, root.Val+right))
}

func maxPathSum(root *TreeNode) int {
	maxSum = math.MinInt32

	findMaxSum(root)
	return maxSum
}

func main() {

}
