package main

import (
	"sort"
)

// 回文数
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
// 可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
func singleNumber(nums []int) int {
	m := make(map[int]int)
	for index := 0; index < len(nums); index++ {
		m[nums[index]]++ //把数组的值放入map的key中，map中value为出现的次数
	}
	for k, v := range m { //循环遍历map中value =1 时，key也就是数组的值
		if v == 1 {
			return k
		}

	}
	return 0
}

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
func isValid(s string) bool {
	stack := []rune{}                                //stack的切片来存放括号
	m := map[rune]rune{')': '(', '}': '{', ']': '['} //创建map中每组括号分别把右括号作为key,左括号作为value
	for _, char := range s {                         //遍历字符串，char是key，val是value
		if val, ok := m[char]; ok { //判断检查char是不是右括号,是否存在
			if len(stack) == 0 || stack[len(stack)-1] != val { //把切片为空或者切片顶部的元素与map中value进行比较，不是左括号
				return false //返回错误
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}

// 写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，返回空字符串
func findLongPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix) && j < len(strs[i]); j++ {
			if prefix[j] != strs[i][j] {
				prefix = prefix[:j]
				break
			}
		}
	}
	return prefix
}

//给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
//将大整数加 1，并返回结果的数字数组。

func sortAndAdd(digits []int) []int {
	n := len(digits) // 获取数组长度2
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}

// 删除数组中的重复元素
func deleteDuplicates(arr []int) int {
	len := len(arr)
	if len == 0 {
		return 0
	}
	slow := 0
	for fast := 1; fast < len; fast++ {
		if arr[fast] != arr[slow] { //11233
			arr[slow+1] = arr[fast]
			slow++
		}
	}

	return slow + 1

}

// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
// 可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
// 将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	// 按照区间的起始位置进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//fmt.Println(intervals)          //排序后打印
	merged := [][]int{intervals[0]} // 初始化合并后的区间数组，包含第一个区间
	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1] // 获取合并后数组中的最后一个区间
		//fmt.Println(last)
		current := intervals[i]
		//fmt.Println(current)
		// 当前遍历的区间
		if current[0] <= last[1] { // 判断是否有重叠
			// 有重叠，更新最后一个区间的结束位置为两者的最大值
			//fmt.Println(current, last)
			last[1] = max(last[1], current[1])
		} else {
			// 没有重叠，直接将当前区间添加到合并后的数组中
			merged = append(merged, current)
		}
	}
	return merged
}

// 计算数组中两数之和等于目标值
func found(arr []int, target int) map[int]int {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				m[arr[i]] = arr[j]
			}
		}
	}

	return m
}
func main() {
	//回文数
	// nums := []int{1, 2, 3, 4, 1, 2, 4}
	// result := singleNumber(nums)
	// fmt.Println(result)
	//括号匹配
	// s := "(()[])}"
	// fmt.Println(isValid(s))

	// strs := []string{"flower", "flow", "flight"}
	// fmt.Println(findLongPrefix(strs))

	// digits := []int{9, 2, 9}
	// fmt.Println(sortAndAdd(digits))

	// arr := []int{1, 1, 2, 3, 3}
	// newLength := deleteDuplicates(arr)
	// fmt.Println(arr[:newLength])

	// 测试合并区间
	// intervals := [][]int{{7, 6}, {1, 3}, {8, 10}, {15, 18}}
	// merged := merge(intervals)
	// fmt.Println(merged)
	// a := []int{1, 3, 5, 7}
	// b := []int{2, 8}
	// c := max(a[len(a)-1], b[len(b)-1])
	// d := min(a[0], b[0])
	// fmt.Println(c)
	// fmt.Println(d)
	// arr := []int{1, 2, 3, 4, 5, 6}
	// target := 5
	// m := found(arr, target)
	// for k, v := range m {
	// 	fmt.Println("Found pair:", k, v)
	// }
}
