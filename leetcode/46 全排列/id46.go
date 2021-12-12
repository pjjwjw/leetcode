package id46

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	n := len(nums)
	used := make([]bool, n)
	rt := [][]int{}
	track := []int{}
	var dfs func()
	dfs = func() {
		if len(track) == n {
			tmp := make([]int, n)
			copy(tmp, track)
			rt = append(rt, tmp)
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			track = append(track, nums[i])
			dfs()
			track = track[0 : len(track)-1]
			used[i] = false
		}
	}
	dfs()
	return rt
}

// @lc code=end
