package main

import "fmt"

/*
题目：最小栈 (Min Stack)
描述：设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
输入示例：
MinStack
push -2
push 0
push -3
getMin
pop
top
getMin
输出示例：
null null null null -3 null 0 -2
*/

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	val := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if val == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func main() {
	obj := Constructor()
	var op string
	for {
		if _, err := fmt.Scan(&op); err != nil {
			break
		}
		if op == "push" {
			var v int
			fmt.Scan(&v)
			obj.Push(v)
			fmt.Print("null ")
		} else if op == "pop" {
			obj.Pop()
			fmt.Print("null ")
		} else if op == "top" {
			fmt.Printf("%d ", obj.Top())
		} else if op == "getMin" {
			fmt.Printf("%d ", obj.GetMin())
		}
	}
	fmt.Println()
}
