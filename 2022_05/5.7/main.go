package main

var (
	cnt [1010][2]int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getPrime(x int) (int, int) {
	cntTwo, cntFive := 0, 0
	for x%2 == 0 {
		cntTwo += 1
		x /= 2
	}
	for x%5 == 0 {
		cntFive += 1
		x /= 5
	}
	return cntTwo, cntFive
}

func init() {
	for i := 1; i < 1010; i++ {
		cnt[i][0], cnt[i][1] = getPrime(i)
	}
}

func maxTrailingZeros(grid [][]int) int {
	n := len(grid)
	if n <= 0 {
		return 0
	}
	m := len(grid[0])

	left := make([][][]int, n)
	right := make([][][]int, n)
	up := make([][][]int, n)
	down := make([][][]int, n)
	for i := 0; i < n; i++ {
		left[i], right[i], up[i], down[i] = make([][]int, m), make([][]int, m), make([][]int, m), make([][]int, m)
		for j := 0; j < m; j++ {
			left[i][j], right[i][j], up[i][j], down[i][j] = make([]int, 2), make([]int, 2), make([]int, 2), make([]int, 2)
		}
	}

	for i := 0; i < n; i++ {
		left[i][0][0], left[i][0][1] = cnt[grid[i][0]][0], cnt[grid[i][0]][1]
		right[i][m-1][0], right[i][m-1][1] = cnt[grid[i][m-1]][0], cnt[grid[i][m-1]][1]
	}
	for j := 0; j < m; j++ {
		up[0][j][0], up[0][j][1] = cnt[grid[0][j]][0], cnt[grid[0][j]][1]
		down[n-1][j][0], down[n-1][j][1] = cnt[grid[n-1][j]][0], cnt[grid[n-1][j]][1]
	}

	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			cntTwo, cntFive := cnt[grid[i][j]][0], cnt[grid[i][j]][1]
			left[i][j][0], left[i][j][1] = left[i][j-1][0]+cntTwo, left[i][j-1][1]+cntFive
			cntTwo, cntFive = cnt[grid[i][m-j-1]][0], cnt[grid[i][m-j-1]][1]
			right[i][m-j-1][0], right[i][m-j-1][1] = right[i][m-j][0]+cntTwo, right[i][m-j][1]+cntFive
		}
	}
	for j := 0; j < m; j++ {
		for i := 1; i < n; i++ {
			cntTwo, cntFive := cnt[grid[i][j]][0], cnt[grid[i][j]][1]
			up[i][j][0], up[i][j][1] = up[i-1][j][0]+cntTwo, up[i-1][j][1]+cntFive
			cntTwo, cntFive = cnt[grid[n-i-1][j]][0], cnt[grid[n-i-1][j]][1]
			down[n-i-1][j][0], down[n-i-1][j][1] = down[n-i][j][0]+cntTwo, down[n-i][j][1]+cntFive
		}
	}

	ret := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			cntTwo, cntFive := getPrime(grid[i][j])
			twoSum := left[i][j][0] + up[i][j][0] - cntTwo
			fiveSum := left[i][j][1] + up[i][j][1] - cntFive
			ret = max(ret, min(twoSum, fiveSum))

			twoSum = left[i][j][0] + down[i][j][0] - cntTwo
			fiveSum = left[i][j][1] + down[i][j][1] - cntFive
			ret = max(ret, min(twoSum, fiveSum))

			twoSum = right[i][j][0] + up[i][j][0] - cntTwo
			fiveSum = right[i][j][1] + up[i][j][1] - cntFive
			ret = max(ret, min(twoSum, fiveSum))

			twoSum = right[i][j][0] + down[i][j][0] - cntTwo
			fiveSum = right[i][j][1] + down[i][j][1] - cntFive
			ret = max(ret, min(twoSum, fiveSum))
		}
	}
	return ret
}
