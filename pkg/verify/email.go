package verify

import "regexp"

// ValidateEmail 校验邮箱格式是否合法
func Email(email string) bool {
	// 正则表达式规则（符合RFC 5322标准）
	// 解释：
	// ^[a-zA-Z0-9._%+-]+  -> 用户名部分（允许字母、数字、._%+-）
	// @[a-zA-Z0-9.-]+     -> @符号及域名部分（允许字母、数字、.-）
	// \.[a-zA-Z]{2,}$     -> 顶级域名（至少2个字母，如.com/.cn）
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// 编译正则表达式
	re := regexp.MustCompile(pattern)

	// 匹配检查
	return re.MatchString(email)
}
