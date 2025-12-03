package utils

import (
	"iter"
	"strings"
)

func SplitByNewLine(s string) iter.Seq[string] {
	return strings.SplitSeq(s, "\r\n")
}

func SplitByComma(s string) iter.Seq[string] {
	return strings.SplitSeq(s, ",")
}
