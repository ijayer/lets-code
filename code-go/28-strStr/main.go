package main

func main() {
	haystack := "hello"
	needle := "ll"

	println(StrStr(haystack, needle))
}

// StrStr
// Solution1: 从原始串中取与字串等长的字符串一词遍比对，直到原始字符串比对完
// 1. calc needle len
// 2. 以 needle 长度为基础，从 haystack 获取等长的字符串 check 是否相等
// 3. 相等时，获取其索引值
func StrStr(haystack, needle string) int {
	if needle == "" {
		return 0
	}

	var ret = -1

	subStrLen := len(needle)
	for i := 0; i <= len(haystack)-subStrLen; i++ {
		target := haystack[i : subStrLen+i]
		if target != needle {
			continue
		}
		ret = i
		break
	}

	return ret
}
