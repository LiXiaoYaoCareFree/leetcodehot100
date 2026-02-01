package main

import (
	"bufio"
	"fmt"
	"os"
)

// 301. 删除无效的括号
// 给你一个由若干括号和字母组成的字符串 s ，删除最小数量的无效括号，使得输入的字符串有效。
// 返回所有可能的结果。答案可以按 任意顺序 返回。

// 示例 1：
// 输入：s = "()())()"
// 输出：["(())()","()()()"]

// 示例 2：
// 输入：s = "(a)())()"
// 输出：["(a())()","(a)()()"]

// 示例 3：
// 输入：s = ")("
// 输出：[""]

func removeInvalidParentheses(s string) []string {
	var res []string
	level := map[string]struct{}{s: {}}

	for len(level) > 0 {
		for s := range level {
			if isValid(s) {
				res = append(res, s)
			}
		}
		if len(res) > 0 {
			return res
		}
		nextLevel := make(map[string]struct{})
		for s := range level {
			for i := 0; i < len(s); i++ {
				if i > 0 && s[i] == s[i-1] {
					continue
				}
				if s[i] == '(' || s[i] == ')' {
					nextLevel[s[:i]+s[i+1:]] = struct{}{}
				}
			}
		}
		level = nextLevel
	}
	return res
}

func isValid(s string) bool {
	cnt := 0
	for _, ch := range s {
		if ch == '(' {
			cnt++
		} else if ch == ')' {
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	return cnt == 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		// Remove quotes if present for easier testing from cli
		if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
			s = s[1 : len(s)-1]
		}
		
		fmt.Println(removeInvalidParentheses(s))
	}
}
