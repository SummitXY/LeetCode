package main
// https://leetcode.cn/problems/generate-parentheses/solution/sui-ran-bu-shi-zui-xiu-de-dan-zhi-shao-n-0yt3/
// 列出所有可能的括号对
import (
	"fmt"
	"strings"
)

func main() {
	res := strings.Join(generateAllParentheses(4), ", ")
	fmt.Println(res)

	res2 := strings.Join(generateSymmetricalParentheses(2), ", ")
	fmt.Println(res2)
}

func generateAllParentheses(n int) []string {
	res := make([]string, 0, 1 << n)
	var dfs func(string)
	dfs = func(p string) {
		if len(p) == n {
			res = append(res, p)
			return
		}

		dfs(p+"(")
		dfs(p+")")
	}

	dfs("")
	return res
}

func generateSymmetricalParentheses(n int) []string {
	res := make([]string, 0, 1 << (2*n))
	var dfs func(string, int, int)
	dfs = func(p string, left, right int) {
		if left > n || right > left {
			return
		}
		if len(p) == 2*n {
			res = append(res, p)
			return
		}

		dfs(p+"(", left+1, right)
		dfs(p+")", left, right+1)
	}

	dfs("", 0, 0)
	return res
}


