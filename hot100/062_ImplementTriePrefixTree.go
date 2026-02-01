package main

import "fmt"

/*
题目：实现 Trie (前缀树) (Implement Trie (Prefix Tree))
描述：Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。
请你实现 Trie 类：
Trie() 初始化前缀树对象。
void insert(String word) 向前缀树中插入字符串 word 。
boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。
输入示例：
Trie
insert apple
search apple
search app
startsWith app
insert app
search app
输出示例：
null null true false true null true
*/

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.SearchPrefix(prefix) != nil
}

func (this *Trie) SearchPrefix(prefix string) *Trie {
	node := this
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func main() {
	obj := Constructor()
	var op string
	for {
		if _, err := fmt.Scan(&op); err != nil {
			break
		}
		if op == "insert" {
			var w string
			fmt.Scan(&w)
			obj.Insert(w)
			fmt.Print("null ")
		} else if op == "search" {
			var w string
			fmt.Scan(&w)
			if obj.Search(w) {
				fmt.Print("true ")
			} else {
				fmt.Print("false ")
			}
		} else if op == "startsWith" {
			var w string
			fmt.Scan(&w)
			if obj.StartsWith(w) {
				fmt.Print("true ")
			} else {
				fmt.Print("false ")
			}
		}
	}
	fmt.Println()
}
