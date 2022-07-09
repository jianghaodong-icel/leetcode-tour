# 题目链接
https://leetcode.com/problems/sequentially-ordinal-rank-tracker/

# 解题思路
这道题可以使用`TreeMap`来解决，保证插入/查询的时间复杂度都是`O(logn)`，之后我们使用一个迭代器指针`cur`，每次查询，返回迭代器指针指向的结果，之后向后移动迭代器指针

需要注意，如果插入的元素在当前迭代器指针指向元素的前面，则需要将该指针向前移动一位，以保证下一次查询的正确性