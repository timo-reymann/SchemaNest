package database

import (
	"strconv"
	"strings"
)

// ReplacePlaceholders replaces each '?' in the input string with $1, $2, ..., $n
func ReplacePlaceholders(query string) string {
	var builder strings.Builder
	argIndex := 1
	inString := false

	for i := 0; i < len(query); i++ {
		if query[i] == '?' && !inString {
			builder.WriteString("$")
			builder.WriteString(strconv.Itoa(argIndex))
			argIndex++
		} else if query[i] == '\'' {
			inString = !inString
			builder.WriteByte(query[i])
		} else {
			builder.WriteByte(query[i])
		}
	}

	return builder.String()
}
