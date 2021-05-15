package main

import "fmt"

func main() {
	fmt.Printf("twoSum: %v\n", TwoSum([]int{2, 7, 11, 15}, 9))
}

// Solution1: 暴力比对
// 	- 时间复杂度：O(n²)
//	- 空间复杂度：O(n)
func TwoSum(nums []int, target int) []int {
	for i, v := range nums {
		for j, vv := range nums[i+1:] {
			if v+vv == target {
				return []int{i, i + 1 + j} // i+1+j 是被加数在 nums 里的索引
			}
		}
	}
	return nil
}

// Solution2：以空间换时间，省去一层循环
// 	- 时间复杂度：O(n)
//	- 空间复杂度：O(n)
func TwoSum2(nums []int, target int) []int {
	m := make(map[int]int) // key:index, val: value

	for i, v := range nums {
		if _, ok := m[target-v]; !ok {
			m[v] = i
			continue
		}
		return []int{m[target-v], i}
	}
	return nil
}
