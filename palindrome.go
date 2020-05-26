package go_algo

import (
	"strings"
)

func Palindrome(s string) bool {
	// 去除字符串中的非字母和数字字符
	filteredStr := filteredString(s)

	// 全部转成小写
	lowerStr := lowerString(filteredStr)

	// 反转字符串
	reversedString := reversedString(lowerStr)

	// 字符串比较
	return strings.EqualFold(lowerStr, reversedString)
}

func filteredString(s string) string {
	return strings.ReplaceAll(s, "[^A-Za-z0-9]", "")
}

func lowerString(s string) string {
	return strings.ToLower(s)
}

func reversedString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}