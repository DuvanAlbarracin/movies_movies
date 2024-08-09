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
