package main

import "sort"

type myOrderedMap struct {
	keys     []int
	innerMap map[int]int
}

func constructMyMap() *myOrderedMap {
	ret := &myOrderedMap{}
	keys := make([]int, 0)
	innerMap := make(map[int]int)
	ret.keys, ret.innerMap = keys, innerMap
	return ret
}

func (m *myOrderedMap) Set(key, val int) {
	_, exists := m.innerMap[key]
	if !exists {
		m.keys = append(m.keys, key)
	}
	m.innerMap[key] = val
}

func (m *myOrderedMap) Get(key int) int {
	val := m.innerMap[key]
	return val
}

func (m *myOrderedMap) SortKeys() {
	sort.Ints(m.keys)
}

func (m *myOrderedMap) GetKeys() []int {
	return m.keys
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type MyCalendarThree struct {
	calendarMap *myOrderedMap
}

func Constructor() MyCalendarThree {
	ret := MyCalendarThree{}
	ret.calendarMap = constructMyMap()
	return ret
}

func (this *MyCalendarThree) Book(start int, end int) int {
	preStartVal := this.calendarMap.Get(start)
	this.calendarMap.Set(start, preStartVal+1)
	preEndVal := this.calendarMap.Get(end)
	this.calendarMap.Set(end, preEndVal-1)

	this.calendarMap.SortKeys()
	kMax, preSum := 0, 0
	for _, timePoint := range this.calendarMap.GetKeys() {
		preSum += this.calendarMap.Get(timePoint)
		kMax = max(kMax, preSum)
	}
	return kMax
}

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
