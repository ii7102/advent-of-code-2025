package utils

import (
	"strconv"
	"strings"
)

func Range(s string) (left, right int) {
	leftStr, rightStr, found := strings.Cut(s, "-")
	if !found {
		panic("invalid string: " + s)
	}

	left, err := strconv.Atoi(leftStr)
	if err != nil {
		panic("invalid left: " + leftStr)
	}

	right, err = strconv.Atoi(rightStr)
	if err != nil {
		panic("invalid right: " + rightStr)
	}

	return
}

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("invalid string: " + s)
	}

	return i
}
