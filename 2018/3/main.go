package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Fabric [1000][1000]int

func main() {
	f := Fabric{}
	for _, l := range readFile("input.txt") {
		s := strings.Split(l, " ")

		offset := strings.Split(s[2][:len(s[2])-1], ",")
		dims := strings.Split(s[3], "x")

		left, _ := strconv.Atoi(offset[0])
		top, _ := strconv.Atoi(offset[1])
		width, _ := strconv.Atoi(dims[0])
		height, _ := strconv.Atoi(dims[1])

		for r := top; r < top+height; r++ {
			for c := left; c < left+width; c++ {
				f[r][c]++
			}
		}
	}
	fmt.Println(f.Overlaps())
	f.print()
}

func (f Fabric) Overlaps() (n int) {
	for i := range f {
		for j := range f[i] {
			if f[i][j] > 1 {
				n++
			}
		}
	}
	return
}

func (f Fabric) print() {
	for i := range f {
		ln := ""
		for j := range f[i] {
			switch f[i][j] {
			case 0:
				ln += "."
			case 1:
				ln += "O"
			default:
				ln += "X"
			}
		}
		fmt.Println(ln)
	}
}

func readFile(fn string) []string {
	dat, _ := ioutil.ReadFile(fn)
	return strings.Split(string(dat), "\n")
}
