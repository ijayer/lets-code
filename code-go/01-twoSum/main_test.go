/*
 * 说明：
 * 作者：zhe
 * 时间：2019-01-25 4:50 PM
 * 更新：
 */

package main

import "testing"

func TestTwoSum(t *testing.T) {
	var nums = []int{2, 7, 11, 15}
	var target = 22
	index := TwoSum(nums, target)
	t.Logf("%v", index)
}
