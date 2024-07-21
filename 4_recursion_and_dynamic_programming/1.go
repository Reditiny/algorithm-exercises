package main

/**
给定整数 N, 返回斐波那契数列的第N项

同：给定整数M 代表台阶数， 一次可以跨 2 个或者 1 个台阶， 返回有多少种走法
*/

// O(2^n)
func fibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciRecursion(n-1) + fibonacciRecursion(n-2)
}

// O(n)
func fibonacciDP(n int) int {
	dp := make([]int, 2)
	dp[0] = 0
	dp[1] = 1
	// dp[i] 在第 i 个台阶的走法，上步是 dp[i-1] 或 dp[i-2]
	// dp[i] = dp[i-1] + dp[i-2]
	for i := 2; i <= n; i++ {
		dp[i%2] = dp[(i-1)%2] + dp[(i-2)%2]
	}
	return dp[n%2]
}

// O(logN)
// F(n) = F(n-1) + F(n-2) 二阶递推数列
// | F(n)   | = | 1 1 | * | F(n-1) | = | 1 1 |^(n-1) * | F(1) |
// | F(n-1) |   | 1 0 |   | F(n-2) |   | 1 0 |         | F(0) |
func fibonacciMatrix(n int) int {
	if n <= 1 {
		return n
	}
	base := [][]int{{1, 1}, {1, 0}}
	res := matrixPower(base, n-1)
	return res[0][0]
}

func matrixPower(base [][]int, n int) [][]int {
	ans := make([][]int, len(base))
	for i := 0; i < len(base); i++ {
		ans[i] = make([]int, len(base[0]))
		ans[i][i] = 1
	}
	tmp := base
	for n != 0 {
		if n&1 == 1 {
			ans = multiMatrix(ans, tmp)
		}
		tmp = multiMatrix(tmp, tmp)
		n >>= 1
	}
	return ans
}

func multiMatrix(matrix1 [][]int, matrix2 [][]int) [][]int {
	res := make([][]int, len(matrix1))
	for i := 0; i < len(matrix1); i++ {
		res[i] = make([]int, len(matrix2[0]))
		for j := 0; j < len(matrix2[0]); j++ {
			for k := 0; k < len(matrix2); k++ {
				res[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}
	return res
}
