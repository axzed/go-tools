package word

import (
	"strings"
	"unicode"
)

// turn words all Uppercase
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// turn words all Lowercase
func ToLower(s string) string {
	return strings.ToLower(s)
}

// turn words to upper camelcase
func UnderscoreToUpperCamelCase(s string) string {
	// 下划线替换为空格
	s = strings.Replace(s, "_", " ", -1)
	// 替换首字符
	s = strings.Title(s)
	// 将空格字符替换为空
	return strings.Replace(s, " ", "", -1)
}

// turn word to lower camelcase
func UnderscoreToLowerCamelCase(s string) string {
	// 复用UnderscoreToUpperCamelCase先进行初步转换
	s = UnderscoreToUpperCamelCase(s)
	// 取出第一位,然后调用unicode.ToLower方法将字符转为小写
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// camelcase to underscore
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			// 首字符小写加入字符串切片
			output = append(output, unicode.ToLower(r))
			continue
		}
		// 如果是大写
		if unicode.IsUpper(r) {
			// 先加入下划线进入切片
			output = append(output, '_')
		}
		// 再将当前的大写字符转为小写加入切片中
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
