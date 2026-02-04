package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
)

// BloomFilter 布隆过滤器结构体
type BloomFilter struct {
	bitMap  []uint64 // 位图，用uint64数组存储，节省空间
	bitSize uint64   // 位图总位数m
	hashNum uint64   // 哈希函数个数k
}

// NewBloomFilter 布隆过滤器构造函数
// n: 预期存储的元素数量
// p: 允许的误判率（0 < p < 1，如0.01表示1%）
func NewBloomFilter(n uint64, p float64) *BloomFilter {
	if n == 0 || p <= 0 || p >= 1 {
		panic("参数无效：n必须大于0，p必须在(0,1)之间")
	}

	// 1. 计算最优位图总位数m = -n * ln(p) / (ln2)^2
	ln2 := math.Log(2)
	m := uint64(-float64(n) * math.Log(p) / (ln2 * ln2))
	// 保证m是64的倍数，适配uint64数组存储（可选，仅为对齐优化）
	if m%64 != 0 {
		m = (m/64 + 1) * 64
	}

	// 2. 计算最优哈希函数个数k = m * ln2 / n
	k := uint64(float64(m) * ln2 / float64(n))
	// 哈希函数数至少为1
	if k < 1 {
		k = 1
	}

	// 初始化位图：m位需要m/64个uint64元素
	bitMap := make([]uint64, m/64)

	return &BloomFilter{
		bitMap:  bitMap,
		bitSize: m,
		hashNum: k,
	}
}

// 核心工具方法：将位图的第bit位设置为1
func (bf *BloomFilter) setBit(bit uint64) {
	// 计算bit位落在哪个uint64元素（索引）
	index := bit / 64
	// 计算bit位在该uint64中的偏移量（0-63）
	offset := bit % 64
	// 按位或置1
	bf.bitMap[index] |= 1 << offset
}

// 核心工具方法：查询位图的第bit位是否为1
func (bf *BloomFilter) getBit(bit uint64) bool {
	index := bit / 64
	offset := bit % 64
	// 按位与判断是否为1
	return bf.bitMap[index]&(1<<offset) != 0
}

// 哈希函数组：对任意字节数组生成k个不同的哈希值（映射到[0, bitSize-1]）
// 选用md5/sha1/sha256三个加密哈希，通过切片拆分生成多个哈希值，保证分布性
func (bf *BloomFilter) hash(data []byte) []uint64 {
	// 生成3种基础哈希值，拼接后拆分出k个哈希值
	h1 := md5.Sum(data)
	h2 := sha1.Sum(data)
	h3 := sha256.Sum256(data)
	// 拼接所有哈希字节，用于生成更多哈希值
	allHash := append(h1[:], append(h2[:], h3[:]...)...)

	// 生成k个哈希值，每个值映射到[0, bitSize-1]
	hashValues := make([]uint64, bf.hashNum)
	for i := uint64(0); i < bf.hashNum; i++ {
		// 每次取8个字节（uint64长度）生成一个哈希值，循环使用allHash
		start := (i * 8) % uint64(len(allHash))
		end := start + 8
		if end > uint64(len(allHash)) {
			end = uint64(len(allHash))
		}
		// 补全8个字节（不足则补0）
		var b [8]byte
		copy(b[:], allHash[start:end])
		// 将字节转为uint64
		hash := binary.BigEndian.Uint64(b[:])
		// 映射到位图范围
		hashValues[i] = hash % bf.bitSize
	}
	return hashValues
}

// Add 向布隆过滤器添加元素（支持[]byte类型，可封装为支持string/int等）
func (bf *BloomFilter) Add(data []byte) {
	hashValues := bf.hash(data)
	for _, bit := range hashValues {
		bf.setBit(bit)
	}
}

// Contains 查询元素是否存在（false=一定不存在，true=可能存在）
func (bf *BloomFilter) Contains(data []byte) bool {
	hashValues := bf.hash(data)
	for _, bit := range hashValues {
		if !bf.getBit(bit) {
			return false
		}
	}
	return true
}

// 封装方法：支持添加string类型元素
func (bf *BloomFilter) AddString(s string) {
	bf.Add([]byte(s))
}

// 封装方法：支持查询string类型元素
func (bf *BloomFilter) ContainsString(s string) bool {
	return bf.Contains([]byte(s))
}

// 测试主函数
func main() {
	// 初始化布隆过滤器：预期存储10000个元素，误判率1%
	bf := NewBloomFilter(10000, 0.01)
	fmt.Printf("布隆过滤器初始化完成：位图大小=%d位，哈希函数数=%d\n", bf.bitSize, bf.hashNum)

	// 添加测试元素
	testElements := []string{"Go语言", "布隆过滤器", "BloomFilter", "Redis", "MySQL"}
	for _, e := range testElements {
		bf.AddString(e)
		fmt.Printf("添加元素：%s\n", e)
	}

	// 测试存在的元素（应返回true，可能存在）
	fmt.Println("\n查询存在的元素：")
	for _, e := range testElements {
		fmt.Printf("%s -> %t\n", e, bf.ContainsString(e))
	}

	// 测试不存在的元素（应返回false，一定不存在）
	fmt.Println("\n查询不存在的元素：")
	notExistElements := []string{"Java", "Python", "RedisCluster", "K8s"}
	for _, e := range notExistElements {
		fmt.Printf("%s -> %t\n", e, bf.ContainsString(e))
	}

	// 测试大量元素后的误判率（可选）
	fmt.Println("\n测试批量元素误判率：")
	total := 10000
	falsePositive := 0
	for i := 0; i < total; i++ {
		elem := fmt.Sprintf("test_%d", i+10000) // 未添加的元素
		if bf.ContainsString(elem) {
			falsePositive++
		}
	}
	fmt.Printf("测试%d个未添加元素，误判数=%d，实际误判率=%.2f%%\n", total, falsePositive, float64(falsePositive)/float64(total)*100)
}
