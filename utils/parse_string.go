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

func Coordinates(s string) (x, y, z int) {
	strSplit := strings.Split(s, ",")
	if len(strSplit) != 3 {
		panic("invalid string: " + s)
	}

	x, err := strconv.Atoi(strSplit[0])
	if err != nil {
		panic("invalid x: " + strSplit[0])
	}

	y, err = strconv.Atoi(strSplit[1])
	if err != nil {
		panic("invalid y: " + strSplit[1])
	}

	z, err = strconv.Atoi(strSplit[2])
	if err != nil {
		panic("invalid y: " + strSplit[2])
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
