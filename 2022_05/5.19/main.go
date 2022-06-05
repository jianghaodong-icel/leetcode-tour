package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func check(num int, quantities []int) int {
	ret := 0
	for _, quantitie := range quantities {
		ret += quantitie/num
		if quantitie%num!=0 {
			ret += 1
		}
	}
	return ret
}

func minimizedMaximum(n int, quantities []int) int {
	leftBound, rightBound := 1, -1
	for i:=0; i<len(quantities); i++ {
		rightBound = max(rightBound, quantities[i])
	}

	l, r := leftBound, rightBound
	for l < r {
		mid := (l + r) >> 1
		if check(mid, quantities) > n {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func main() {
}
