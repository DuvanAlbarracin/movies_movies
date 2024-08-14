package utils

import (
	"fmt"
	"html/template"
	"strings"
)

func TrimString(s string) string {
	return strings.TrimRight(s, "\r\n")
}

func SanitizeString(s string) string {
	return template.HTMLEscapeString(s)
}

func BuildUpdateQueryString(table string, id int64, fields map[string]any) (string, error) {
	var b strings.Builder
	_, err := b.WriteString("UPDATE ")
	if err != nil {
		return "", err
	}

	fmt.Fprintf(&b, "%s SET", table)

	index := 1
	for key := range fields {
		fmt.Fprintf(&b, " %s = $%d", key, index)
		if index != len(fields) {
			b.WriteRune(',')
		}
		index++
	}

	fmt.Fprintf(&b, " WHERE id = %d;", id)

	return b.String(), nil
}

func ToSnakeCase(s string) string {
	var sb strings.Builder
	offset := 0

	if isUpper(s[0]) {
		offset = 32
	}
	sb.WriteByte(s[0] + byte(offset))

	for i := 1; i < len(s); i++ {
		offset = 0
		if isUpper(s[i]) {
			offset = 32
			sb.WriteRune('_')
		}
		sb.WriteByte(s[i] + byte(offset))
	}

	return sb.String()
}

func isUpper(b byte) bool {
	if int(b) > 65 && int(b) < 90 {
		return true
	}
	return false
}
