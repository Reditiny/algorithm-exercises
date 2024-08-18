package main

/**
给定一个字符串数组 strs， 其中有些位置为 null, 在不为 null 的位置上，其字符串是按照字典顺序由小到大依次出现的。
再给定一个字符串 str, 请返回 str 在 strs 中出现的最左的位置
*/

func findStrIndex(arr []string, target string) int {
	right := len(arr)
	if right == 0 {
		return -1
	}
	ans := 0

	left := 0
	for left < right {
		mid := (right + left) / 2
		if arr[mid] == "" {
			find := false
			for temp := mid - 1; temp >= left; temp-- {
				if arr[temp] == "" {
					continue
				}
				if arr[temp] == target {
					ans = temp
					right = temp - 1
					find = true
				} else if arr[temp] > target {
					right = temp - 1
					find = true
				}
				break
			}
			if !find {
				for temp := mid + 1; temp <= right; temp++ {
					if arr[temp] == "" {
						continue
					}
					if arr[temp] == target {
						return temp
					}
					left = temp + 1
					break
				}
			}
		} else if arr[mid] == target {
			ans = mid
			right = mid - 1
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}
