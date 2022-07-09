package main

func grayCode(n int) []int {
	ret := make([]int, 0)
	for i := 0; i < 1<<n; i++ {
		ret = append(ret, i^(i>>1))
	}
	return ret
}

func main() {

}
