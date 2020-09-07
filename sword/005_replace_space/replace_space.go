package _05_replace_space

import "strings"

//请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
//
//
//
//示例 1：
//
//输入：s = "We are happy."
//输出："We%20are%20happy."
//
//
//限制：
//
//0 <= s 的长度 <= 10000
//

func replaceSpace(s string) string {
	result := make([]rune, len(s)*3)

	i := 0
	for _, v := range s {
		if v != ' ' {
			result[i] = v
			i++
		} else {
			result[i] = '%'
			result[i+1] = '2'
			result[i+2] = '0'
			i += 3

		}
	}

	return string(result)[:i]

}

func replaceSpace1(s string) string {
	var res strings.Builder
	for i := range s {
		if s[i] == ' ' {
			res.WriteString("%20")
		} else {
			res.WriteByte(s[i])
		}
	}
	return res.String()
}
