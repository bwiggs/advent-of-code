package lib

import (
	"io/ioutil"
	"strings"
)

func ReadLines(fn string) []string {
	dat, _ := ioutil.ReadFile(fn)
	return strings.Split(string(dat), "\n")
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
