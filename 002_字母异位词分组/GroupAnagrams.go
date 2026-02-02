package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	strs := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&strs[i])
	}
	res := groupAnagrams(strs)
	fmt.Println(res)
}

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, v := range strs {
		tmp := []byte(v)
		slices.Sort(tmp)
		SortedS := string(tmp)
		m[SortedS] = append(m[SortedS], v)
	}
	return slices.Collect(maps.Values(m))
}

/*

@'
6
eat tea tan ate nat bat
'@ | go run .\GroupAnagrams.go

*/
