package id47

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	n := len(nums)
	use := make([]bool, n)
	rt := [][]int{}
	track := []int{}
	var dfs func()
	dfs = func() {
		if len(track) == n {
			tmp := make([]int, n)
			copy(tmp, track)
			rt = append(rt, tmp)
		}
		has := map[int]bool{}
		for i := 0; i < n; i++ {
			if use[i] || has[nums[i]] {
				continue
			}
			has[nums[i]] = true
			use[i] = true
			track = append(track, nums[i])
			dfs()
			track = track[0 : len(track)-1]
			use[i] = false
		}
	}
	dfs()
	return rt
}

// @lc code=end
