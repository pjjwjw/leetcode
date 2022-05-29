package id49

import (
	"strconv"
)

func groupAnagrams(strs []string) [][]string {
	return nil
}

func encode1(str string) string {
	m := [26]int{}
	for _, c := range str {
		m[c-'0']++
	}

	rt := ""
	for i, v := range m {
		if v != 0 {
			rt += strconv.Itoa(i+'0') + strconv.Itoa(v)
		}
	}

	return rt
}
