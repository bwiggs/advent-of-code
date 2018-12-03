package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	boxIDs := readBoxes()
	fmt.Printf("checksum: %d\n", checksum(boxIDs))
	fmt.Printf("box characters: %s\n", findFabric(boxIDs))
}

func findFabric(boxIDs []string) string {
	for i := 0; i < len(boxIDs)-1; i++ {
		for j := i + 1; j < len(boxIDs); j++ {
			if found, dIdx := strdist(boxIDs[i], boxIDs[j]); found {
				return string(boxIDs[i][:dIdx]) + string(boxIDs[i][dIdx+1:])
			}
		}
	}
	return ""
}

func strdist(a, b string) (bool, int) {
	// fmt.Println(a + "\n" + b + "\n")
	var dist int
	di := -1
	for i := range a {
		if a[i] != b[i] {
			dist++
			di = i
		}
	}
	return dist == 1, di
}

func checksum(boxIDs []string) int {
	var twos, threes int
	for _, bid := range boxIDs {
		chars := make(map[rune]int)
		for _, ch := range bid {
			if _, ok := chars[ch]; ok {
				chars[ch]++
			} else {
				chars[ch] = 1
			}
		}

		var f2, f3 bool
		for k := range chars {
			if !f2 && chars[k] == 2 {
				twos++
				f2 = true
			}

			if !f3 && chars[k] == 3 {
				threes++
				f3 = true
			}
		}
	}
	return twos * threes
}

func readBoxes() []string {
	dat, _ := ioutil.ReadFile("input.txt")
	return strings.Split(string(dat), "\n")
}
