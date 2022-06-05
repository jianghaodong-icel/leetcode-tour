package main

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	MIN := -100000

	ret := make([][]int, 0)
	if n == 0 {
		return ret
	}

	// 将区间按照左端点从小到大排序
	sort.Slice(intervals, func(i int, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	l, r := MIN, MIN
	for i := 0; i < n; i++ {
		if intervals[i][0] > r {
			if l != MIN {
				ret = append(ret, []int{l, r})
			}
			l = intervals[i][0]
			r = intervals[i][1]
		} else {
			r = max(r, intervals[i][1])
		}
	}
	if l != MIN {
		ret = append(ret, []int{l, r})
	}
	return ret
}

func main() {

}
