package main

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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func check(arr []int, value int) int {
	ret := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] >= value {
			ret += value
		} else {
			ret += arr[i]
		}
	}
	return ret
}

func findBestValue(arr []int, target int) int {
	n := len(arr)

	leftBound, rightBound := 0, 0
	for i := 0; i < n; i++ {
		rightBound = max(rightBound, arr[i])
	}

	minDiff, ret := 0x3f3f3f3f, 0x3f3f3f3f
	l, r := leftBound, rightBound
	for l <= r {
		mid := (l + r) >> 1
		cTarget := check(arr, mid)
		if cTarget > target {
			r = mid - 1
		} else {
			l = mid + 1
		}

		absDiff := abs(cTarget - target)
		if absDiff < minDiff {
			ret = mid
			minDiff = absDiff
		} else if absDiff == minDiff {
			ret = min(ret, mid)
		}
	}
	return ret
}

func main() {

}
