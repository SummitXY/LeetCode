package main

import (
	"fmt"
)

// https://leetcode.cn/problems/permutations/solution/hui-su-suan-fa-python-dai-ma-java-dai-ma-by-liweiw/
// 列出所有的数字排序组合
func main() {

	fmt.Println(generateAllPermute([]int{1,2,3}))
}


func generateAllPermute(nums []int) [][]int {
	res := make([][]int, 0, 100)
	visited := make([]int, len(nums))
	tmp := make([]int, 0, len(nums))

	var dfs func()
	dfs = func() {
		if len(tmp) == len(nums) {
			ans := make([]int, len(tmp))
			copy(ans, tmp)
			res = append(res, ans)
			return
		}
		for i:=0; i<len(nums); i++ {
			if visited[i] == 1 {
				continue
			}
			tmp = append(tmp, nums[i])
			visited[i]=1
			dfs()
			visited[i]=0
			tmp = tmp[:len(tmp)-1]
		}
	}

	dfs()
	return res
}


