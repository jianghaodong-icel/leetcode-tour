package main

func carPooling(trips [][]int, capacity int) bool {
	N := 1010
	preSum := make([]int, N)
	for _, trip := range trips {
		numPassengers := trip[0]
		l, r := trip[1]+1, trip[2]+1
		preSum[l] += numPassengers
		preSum[r] -= numPassengers
	}

	for i:=1; i<N; i++ {
		preSum[i] += preSum[i-1]
		if preSum[i] > capacity {
			return false
		}
	}
	return true
}

func main() {

}
