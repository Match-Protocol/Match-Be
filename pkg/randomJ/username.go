package randomJ

import (
	"github.com/kpechenenko/rword"
	"math/rand"
	"time"
)

func GenerateRandomEnglishName() string {
	names := []string{
		"James", "John", "Michael", "David", "William",
		"Mary", "Jennifer", "Lisa", "Susan", "Elizabeth",
		"Alex", "Taylor", "Jordan", "Casey", "Riley",
	}

	rand.Seed(time.Now().UnixNano())
	return names[rand.Intn(len(names))]
}

func GenerateRandomEnglishNameV2() string {
	g, err := rword.New() // 默认加载内置词库
	if err != nil {
		return GenerateRandomEnglishName()
	}
	return g.Word() // 随机单词
}
