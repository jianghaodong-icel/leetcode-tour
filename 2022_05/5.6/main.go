package main

var (
	maxPath = 1
	tree    map[int][]int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxPath(root int, s string) int {
	idx := root
	if idx == -1 {
		idx = 0
	}

	first, second := 0, 0 // 存储最大值 & 次大值
	for _, son := range tree[root] {
		cur := findMaxPath(son, s)

		if s[idx] == s[son] {
			continue
		}
		if cur > first {
			second = first
			first = cur
		} else if cur > second {
			second = cur
		}
	}
	maxPath = max(maxPath, 1+first+second)
	return 1 + first
}

func longestPath(parent []int, s string) int {
	n := len(parent)
	maxPath = 1
	root := -1
	tree = make(map[int][]int)

	for i := 0; i < n; i++ {
		if parent[i] == -1 {
			root = i
			continue
		}
		tree[parent[i]] = append(tree[parent[i]], i)
	}

	findMaxPath(root, s)
	return maxPath
}

func main() {

}