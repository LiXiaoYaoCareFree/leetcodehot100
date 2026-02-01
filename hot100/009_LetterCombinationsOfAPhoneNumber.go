package main

import "fmt"

/*
题目：电话号码的字母组合 (Letter Combinations of a Phone Number)
描述：给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
输入示例：
23
输出示例：
[ad ae af bd be bf cd ce cf]
*/

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		digit := string(digits[index])
		letters := phoneMap[digit]
		for i := 0; i < len(letters); i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
}

func main() {
	var digits string
	fmt.Scan(&digits)
	res := letterCombinations(digits)
	fmt.Println(res)
}
