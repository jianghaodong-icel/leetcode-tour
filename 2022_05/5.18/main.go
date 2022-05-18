package main

import "sort"

func maxSumRangeQuery(nums []int, requests [][]int) int {
	n := len(nums)
	MOD := 1000000007
	cnt := make([]int, n+1)
	for _, request := range requests {
		cnt[request[0]] += 1
		cnt[request[1]+1] -= 1
	}
	for i := 1; i < n; i++ {
		cnt[i] += cnt[i-1]
	}

	sort.Ints(nums)
	sort.Ints(cnt)
	ret := 0
	for i := n - 1; i >= 0; i-- {
		if cnt[i+1] == 0 {
			break
		}
		ret = (ret + (cnt[i+1]*nums[i])%MOD) % MOD
	}
	return ret
}

func main() {

}
