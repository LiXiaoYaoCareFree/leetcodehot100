package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s)
	fmt.Scan(&t)
	fmt.Println(minWindow(s, t))
}

func isCovered(cntS, cntT []int) bool {
	for i := 'A'; i <= 'Z'; i++ {
		if cntS[i] < cntT[i] {
			return false
		}
	}
	for i := 'a'; i <= 'z'; i++ {
		if cntS[i] < cntT[i] {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {
	var cntS, cntT [128]int
	for _, c := range t {
		cntT[c]++
	}

	ansLeft, ansRight := -1, len(s)
	left := 0
	for right, c := range s {
		cntS[c]++
		for isCovered(cntS[:], cntT[:]) {
			if right-left < ansRight-ansLeft {
				ansLeft, ansRight = left, right
			}
			cntS[s[left]]--
			left++
		}
	}
	if ansLeft < 0 {
		return ""
	}
	return s[ansLeft : ansRight+1]
}
