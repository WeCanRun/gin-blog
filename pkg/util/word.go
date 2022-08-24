package util

import "unicode"

// 驼峰 => 下划线
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for k, v := range s {
		if k == 0 {
			output = append(output, unicode.ToLower(v))
			continue
		}

		if unicode.IsUpper(v) {
			output = append(output, '_')
		}

		output = append(output, unicode.ToLower(v))

	}

	return string(output)
}
