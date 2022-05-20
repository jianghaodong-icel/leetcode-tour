package main

import "sort"

type myOrderedMap struct {
	innerMap map[int]int
	keys     []int
}

func New() *myOrderedMap {
	return &myOrderedMap{
		innerMap: make(map[int]int),
		keys:     make([]int, 0),
	}
}

func (m *myOrderedMap) Set(key, val int) {
	_, exists := m.innerMap[key]
	if !exists {
		m.keys = append(m.keys, key)
	}
	m.innerMap[key] = val
}

func (m *myOrderedMap) Get(key int) int {
	return m.innerMap[key]
}

func (m *myOrderedMap) Keys() []int {
	return m.keys
}

func (m *myOrderedMap) SortKeys() {
	sort.Ints(m.keys)
}

func splitPainting(segments [][]int) [][]int64 {
	orderedMap := New()
	for _, segment := range segments {
		preLeftNum := orderedMap.Get(segment[0])
		orderedMap.Set(segment[0], preLeftNum+segment[2])

		preRightNum := orderedMap.Get(segment[1])
		orderedMap.Set(segment[1], preRightNum-segment[2])
	}

	ret := make([][]int64, 0)
	orderedMap.SortKeys()
	preSum := 0
	preKey := 0

	for _, key := range orderedMap.Keys() {
		if preSum > 0 {
			ret = append(ret, []int64{int64(preKey), int64(key), int64(preSum)})
		}
		preSum += orderedMap.Get(key)
		preKey = key
	}
	return ret
}

func main() {

}
