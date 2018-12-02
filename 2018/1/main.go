package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	drifts := readInput()

	freq := 0
	for _, d := range drifts {
		freq += d
	}
	fmt.Printf("Freq: %d\n", freq)

	vf := make(map[int]bool)
	vf[0] = true
	var cf, rf int
	found := false
driftloop:
	for _, d := range drifts {
		cf += d
		if _, ok := vf[cf]; ok {
			rf = cf
			found = true
			break
		}
		vf[cf] = true
	}
	if !found {
		goto driftloop
	}
	fmt.Printf("Freq Rep: %d\n", rf)
}

func readInput() []int {
	fh, _ := os.Open("input.txt")
	defer fh.Close()
	drifts := []int{}
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		d, _ := strconv.Atoi(scanner.Text())
		drifts = append(drifts, d)
	}
	fmt.Println(drifts)
	return drifts
}
