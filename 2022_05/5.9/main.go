package main

import (
	"fmt"
	"sort"
)

func checkValid(x, y, n, m int) bool {
	return (x >= 0 && x < n) && (y >= 0 && y < m)
}

func getBiggestThree(grid [][]int) []int {
	n := len(grid)
	if n <= 0 {
		return []int{}
	}
	m := len(grid[0])

	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k <= j; k++ {
				leftX, leftY := i+k, j-k
				if !checkValid(leftX, leftY, n, m) {
					continue
				}
				rightX, rightY := i+k, j+k
				if !checkValid(rightX, rightY, n, m) {
					continue
				}
				downX, downY := i+2*k, j
				if !checkValid(downX, downY, n, m) {
					continue
				}

				if downX == i && downY == j {
					nums = append(nums, grid[i][j])
					continue
				}

				rhombusSum := 0
				preLeftY, preRightY := leftY, rightY
				for u := leftX; u > i; u-- {
					rhombusSum += grid[u][leftY]
					leftY += 1

					rhombusSum += grid[u][rightY]
					rightY -= 1
				}
				rhombusSum += grid[i][j]

				leftY, rightY = preLeftY+1, preRightY-1
				for u := leftX + 1; u < downX; u++ {
					rhombusSum += grid[u][leftY]
					leftY += 1

					rhombusSum += grid[u][rightY]
					rightY -= 1
				}
				rhombusSum += grid[downX][downY]

				nums = append(nums, rhombusSum)
			}
		}
	}
	sort.Ints(nums)
	ret := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		if i == 0 || nums[i] != nums[i-1] {
			ret = append(ret, nums[i])
		}
	}
	if len(ret) < 3 {
		return ret
	}
	return ret[:3]
}

func main() {
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(getBiggestThree(grid))
}
