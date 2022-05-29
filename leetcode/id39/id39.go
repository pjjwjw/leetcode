package id39

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	rt := [][]int{}
	n := len(candidates)
	track := []int{}
	var dfs func(int, int)
	dfs = func(i int, target int) {
		if i == n || target < 0 {
			return
		}
		if target == 0 {
			tmp := make([]int, len(track))
			copy(tmp, track)
			rt = append(rt, tmp)
			return
		}
		track = append(track, candidates[i])
		dfs(i, target-candidates[i])
		track = track[0 : len(track)-1]
		dfs(i+1, target)
	}
	dfs(0, target)
	return rt
}

// @lc code=end
