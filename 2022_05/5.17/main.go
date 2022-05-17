package main

func canReach(s string, minJump int, maxJump int) bool {
	n := len(s)
	if s[n-1] == '1' {
		return false
	}

	arr := make([]int, n+1)
	arr[minJump], arr[maxJump+1] = 1, -1

	preSum := 0
	for i := 0; i < n; i++ {
		preSum += arr[i]
		if preSum == 0 {
			continue
		}
		if s[i] == '1' {
			continue
		}

		if i+minJump <= n {
			arr[i+minJump] += 1
		}
		if i+maxJump+1 <= n {
			arr[i+maxJump+1] -= 1
		}
	}
	return preSum > 0
}

func main() {

}
