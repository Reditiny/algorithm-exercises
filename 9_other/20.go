package main

/**
升序数组 arr 可能经过一次旋转处理，也可能没有，且 arr 可能存在重复的数。
例如，有序数组［1,2,3,4,5,6,7] 可以旋转处理成［4,5,6,7,1,2,3］
给定一个可能旋转过的有序数组 arr 和一个数 num, 返回 arr 中是否含有 num
*/

func findTarget(nums []int, target int) bool {
	n := len(nums)
	if n == 1 && nums[0] == target {
		return true
	}
	left, right := 0, n-1
	for left < right {
		if nums[left] == target || nums[right] == target {
			return true
		}
		for left < right && nums[left] == nums[right] {
			left++
		}
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		} else if nums[mid] > target && target > nums[left] {
			right = mid - 1
		} else if nums[mid] < nums[left] && (target < nums[mid] || target > nums[left]) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
