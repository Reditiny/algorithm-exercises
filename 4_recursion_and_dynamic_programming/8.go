package main

import (
	ds "algorithm-exercises/0_data_structure"
	"fmt"
)

/**
给定两个字符串 strl 和 str2, 返回两个字符串的最长公共子串
*/

func commonSubStr(str1, str2 string) string {

	dp, maxVal, last := getCommonSubStrDp(str1, str2)
	ans := make([]byte, maxVal)
	fmt.Println(dp)
	for i := maxVal - 1; i >= 0; i-- {
		ans[i] = str1[last]
		last--
	}
	return string(ans)
}

// 获得最长公共子串的 dp 数组
func getCommonSubStrDp(str1, str2 string) ([][]int, int, int) {
	n := len(str1)
	m := len(str2)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	// dp[i][j] 以第 i 和第 j 个元素结尾的最长公共子串长度
	// dp[i][j] = 0 if str1[i] != str2[j]
	// dp[i][j] = dp[i-1][j-1] + 1 if str1[i] == str2[j]
	maxVal := 0
	last := 0

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if str1[i-1] == str2[j-1] {
				// 此处只需要自己左上角的数据，因此可以将空间优化至 O(1)
				dp[i][j] = dp[i-1][j-1] + 1
				last = i - 1
			}
			maxVal = ds.Max(maxVal, dp[i][j])
		}
	}

	return dp, maxVal, last
}
