package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

// 461. 汉明距离
// 两个整数之间的 汉明距离 指的是这两个数字对应二进制位不同的位置的数目。
// 给你两个整数 x 和 y，计算并返回它们之间的汉明距离。

// 示例 1：
// 输入：x = 1, y = 4
// 输出：2
// 解释：
// 1   (0 0 0 1)
// 4   (0 1 0 0)
//        ↑   ↑
// 上面的箭头指出了对应二进制位不同的位置。

// 示例 2：
// 输入：x = 3, y = 1
// 输出：1

func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			fmt.Println(hammingDistance(x, y))
		}
	}
}
