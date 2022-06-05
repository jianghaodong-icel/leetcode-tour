package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDifference(nums []int, queries [][]int) []int {
	n := len(nums)

	numCnt := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		numCnt[i] = make([]int, 110)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= 100; j++ {
			if j == nums[i-1] {
				numCnt[i][j] = numCnt[i-1][j] + 1
			} else {
				numCnt[i][j] = numCnt[i-1][j]
			}
		}
	}

	ret := make([]int, 0, len(queries))
	for _, query := range queries {
		l, r := query[0], query[1]+1
		preNum := -1
		minDif := 0x3f3f3f3f
		for i := 1; i <= 100; i++ {
			cnt := numCnt[r][i] - numCnt[l][i]
			if cnt > 0 {
				if preNum == -1 {
					preNum = i
				} else {
					minDif = min(minDif, i-preNum)
					preNum = i
				}
			}
		}
		if minDif == 0x3f3f3f3f {
			ret = append(ret, -1)
		} else {
			ret = append(ret, minDif)
		}

	}
	return ret
}

func main() {
	nums := []int{4, 5, 2, 2, 7, 10}
	queries := [][]int{
		{2, 3},
		{0, 2},
		{0, 5},
		{3, 5},
	}
	fmt.Println(minDifference(nums, queries))
}
