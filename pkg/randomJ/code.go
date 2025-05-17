package randomJ

import (
	"math/rand"
	"strconv"
	"time"
)

func Number(digits int) string {
	if digits <= 0 {

	}

	// 设置随机种子（确保每次运行结果不同）
	rand.Seed(time.Now().UnixNano())

	// 计算最小值和最大值
	min := intPow(10, digits-1)   // 如3位数: 10^(3-1) = 100
	max := intPow(10, digits) - 1 // 如3位数: 10^3 -1 = 999

	// 生成区间内的随机数
	return strconv.Itoa(rand.Intn(max-min+1) + min)
}

func intPow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}
