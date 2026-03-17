package lib

import "strings"

func findLineAndCol(input string, offset int64) (int, int) {
	line := strings.Count(input[:offset], "\n") + 1
	lastNewLine := strings.LastIndex(input[:offset], "\n")
	column := int(offset) - lastNewLine
	return line, column
}
