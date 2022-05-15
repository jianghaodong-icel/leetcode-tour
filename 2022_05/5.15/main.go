package main

func corpFlightBookings(bookings [][]int, n int) []int {
	preSum := make([]int, n+2)
	for _, booking := range bookings {
		preSum[booking[0]] += booking[2]
		preSum[booking[1]+1] -= booking[2]
	}
	for i:=1; i<=n; i++ {
		preSum[i] += preSum[i-1]
	}

	ret := make([]int, 0, n)
	for i:=1; i<=n; i++ {
		ret = append(ret, preSum[i])
	}
	return ret
}
