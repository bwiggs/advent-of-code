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

// Abs absolute values between 2 ints
func Abs(a int) int {
	if a < 0 {
		return 0 - a
	}
	return a
}

// RectlinearDist returns the manhattan, aka taxi cab distance between 2 points.
func RectlinearDist(ax, ay, bx, by int) int {
	return Abs(ax-bx) + Abs(ay-by)
}
