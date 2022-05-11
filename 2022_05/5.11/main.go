package main

import (
	"fmt"
	"math"
)

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func generatePalindromeNum(x int, isEven bool) int {
	numArr := make([]int, 0)
	for x > 0 {
		numArr = append(numArr, x%10)
		x /= 10
	}

	palindromeNumArr := make([]int, 2*len(numArr))
	if !isEven {
		palindromeNumArr = make([]int, 2*len(numArr)-1)
	}
	for i := 0; i < len(numArr); i++ {
		palindromeNumArr[i] = numArr[len(numArr)-i-1]
	}
	for i := len(palindromeNumArr) - 1; i >= len(palindromeNumArr)/2; i-- {
		palindromeNumArr[i] = palindromeNumArr[len(palindromeNumArr)-i-1]
	}

	ret := 0
	for i := 0; i < len(palindromeNumArr); i++ {
		ret = ret*10 + palindromeNumArr[i]
	}
	return ret
}

// generatePalindromeNumOptimize 优化获取十进制数的方式，不需要使用数组
func generatePalindromeNumOptimize(x int, isEven bool) int {
	reverseNum := 0
	y := x
	if !isEven {
		y /= 10
	}

	cnt := 0
	for y > 0 {
		reverseNum = reverseNum*10 + y%10
		y /= 10
		cnt += 1
	}
	return pow(10, cnt)*x + reverseNum
}

func check(palindromeNum int, k int) bool {
	ret := make([]int, 0)
	for palindromeNum > 0 {
		ret = append(ret, palindromeNum%k)
		palindromeNum /= k
	}

	l, r := 0, len(ret)-1
	for l < r {
		if ret[l] != ret[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func kMirror(k int, n int) int64 {
	len := 1 // 表示当前枚举数字长度

	cnt, ret := 0, 0
	for {
		for i := pow(10, len-1); i <= pow(10, len); i++ {
			palindromeNum := generatePalindromeNumOptimize(i, false)
			if check(palindromeNum, k) {
				cnt += 1
				ret += palindromeNum
			}
			if cnt == n {
				return int64(ret)
			}
		}

		for i := pow(10, len-1); i <= pow(10, len); i++ {
			palindromeNum := generatePalindromeNumOptimize(i, true)
			if check(palindromeNum, k) {
				cnt += 1
				ret += palindromeNum
			}
			if cnt == n {
				return int64(ret)
			}
		}
		len++
	}
}

func main() {
	n, k := 17, 7
	fmt.Println(kMirror(k, n))
}
