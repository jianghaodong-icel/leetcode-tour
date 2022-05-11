package main

func isCovered(ranges [][]int, left int, right int) bool {
	cnt := make([]int, 55)
	for _, rangeArr := range ranges {
		for i := rangeArr[0]; i <= rangeArr[1]; i++ {
			cnt[i] += 1
		}
	}

	for i := left; i <= right; i++ {
		if cnt[i] == 0 {
			return false
		}
	}
	return true
}

func main() {

}
