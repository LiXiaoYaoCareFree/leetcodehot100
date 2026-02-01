package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		if scanner.Scan() {
			p := scanner.Text()
			fmt.Println(findAnagrams(s, p))
		}
	}
}

func findAnagrams(s string, p string) (ans []int) {
	cntP := [26]int{}
	cntS := [26]int{}
	n := len(p)
	for _, v := range p {
		cntP[v-'a']++
	}
	for right, v := range s {
		cntS[v-'a']++
		left := right - n + 1
		if left < 0 {
			continue
		}
		if cntS == cntP {
			ans = append(ans, left)
		}
		cntS[s[left]-'a']--
	}
	return
}
