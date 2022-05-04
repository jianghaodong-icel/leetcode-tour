package main

import "fmt"

var (
	maxNum = 0           // 存储最终的结果
	total  int           // 存储树中节点的总数
	node   []int         // 存储每个节点对应的结果
	tree   map[int][]int // 存储树
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxNum(rootNode int) int {
	partTotal := 1
	partMulti := 1

	for _, sonNode := range tree[rootNode] {
		curNum := findMaxNum(sonNode)
		partTotal += curNum
		partMulti *= curNum
	}
	if total == partTotal {
		node[rootNode] = partMulti
	} else {
		node[rootNode] = partMulti * (total - partTotal)
	}
	maxNum = max(maxNum, node[rootNode])
	return partTotal
}

func countHighestScoreNodes(parents []int) int {
	n := len(parents)
	total = n
	maxNum = 0
	node = make([]int, n)
	tree = make(map[int][]int)

	root := -1
	for i := 0; i < n; i++ {
		if parents[i] == -1 {
			root = i
			continue
		}
		tree[parents[i]] = append(tree[parents[i]], i)
	}
	findMaxNum(root)

	ret := 0
	for i := 0; i < n; i++ {
		if node[i] == maxNum {
			ret++
		}
	}
	return ret
}

func main() {
	parents := []int{-1, 2, 0}
	fmt.Println(countHighestScoreNodes(parents))
}
