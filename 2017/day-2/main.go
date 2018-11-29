package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// The spreadsheet consists of rows of apparently-random numbers. To make sure the recovery process is on the right track, they need you to calculate the spreadsheet's checksum. For each row, determine the difference between the largest value and the smallest value; the checksum is the sum of all of these differences.

// For example, given the following spreadsheet:

// 5 1 9 5
// 7 5 3
// 2 4 6 8
// The first row's largest and smallest values are 9 and 1, and their difference is 8.
// The second row's largest and smallest values are 7 and 3, and their difference is 4.
// The third row's difference is 6.
// In this example, the spreadsheet's checksum would be 8 + 4 + 6 = 18.

// What is the checksum for the spreadsheet in your puzzle input?

func main() {

	var dig string
	min, max, chkSum := -1, -1, 0

	// f := []byte("5\t1\t9\t5\n7\t5\t3\n2\t4\t6\t8\n")
	// scanner := bufio.NewScanner(bytes.NewReader(f))
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {
		t := scanner.Text()
		// spew.Dump(t)
		switch t {
		case "\t": // at end of digit, process
			d, _ := strconv.Atoi(dig)
			if min == -1 {
				min = d
				max = d
			} else if d < min {
				min = d
			} else if d > max {
				max = d
			}
			// fmt.Printf("min: %d max:%d\n", min, max)
			dig = ""
		case "\n":
			d, _ := strconv.Atoi(dig)
			if min == -1 {
				min = d
				max = d
			} else if d < min {
				min = d
			} else if d > max {
				max = d
			}
			// fmt.Printf("min: %d max:%d\n", min, max)
			dig = ""

			chkSum += max - min
			// fmt.Printf("checksum: row: %d file:%d\n", max-min, chkSum)
			min = -1
			max = -1
		default:
			dig += t
		}
	}
	fmt.Printf("part 1 checksum: %d\n", chkSum)
	part2()
}

func part2() {
	f, _ := os.Open("input.txt")
	b, _ := ioutil.ReadAll(f)
	lines := strings.Split(string(b), "\n")
	spreadsheet := make([][]int, len(lines))
	for ri, l := range lines {
		values := strings.Split(l, "\t")
		spreadsheet[ri] = make([]int, len(values))
		for ci, v := range values {
			intVal, _ := strconv.Atoi(v)
			spreadsheet[ri][ci] = intVal
		}
	}

	ss := spreadsheet

	chkSum := 0
	for r := range ss {
		res := 0
		for i := range ss[r] {

			if ss[r][i] == 0 {
				continue
			}
			found := false
			for j := range ss[r] {
				if ss[r][j] == 0 {
					continue
				}
				if i != j && ss[r][i] >= ss[r][j] && ss[r][i]%ss[r][j] == 0 {
					// fmt.Printf("found it! %d/%d=%d\n", ss[r][i], ss[r][j], ss[r][i]/ss[r][j])
					res = ss[r][i] / ss[r][j]
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		chkSum += res
	}
	fmt.Printf("part 2 checksum: %d\n", chkSum)
}
