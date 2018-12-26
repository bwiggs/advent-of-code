package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bwiggs/advent-of-code/2018/lib"
)

func main() {
	edges := make(map[string][]string)
	degrees := make(map[string]int)
	for _, l := range lib.ReadLines("input.txt") {
		words := strings.Split(l, " ")
		succ := words[1]
		dep := words[7]

		if _, ok := edges[succ]; !ok {
			edges[succ] = []string{dep}
		} else {
			edges[succ] = append(edges[succ], dep)
		}

		if _, ok := degrees[dep]; !ok {
			degrees[dep] = 1
		} else {
			degrees[dep]++
		}

		if _, ok := degrees[succ]; !ok {
			degrees[succ] = 0
		}
	}

	var q []string
	ans := []string{}
	for {

		// queue up all ready tasks
		q = []string{}
		for step, deg := range degrees {
			if deg == 0 {
				q = append(q, step)
			}
		}

		// if there's no tasks wer'e done
		if len(q) == 0 {
			break
		}

		// sort the steps in alphabetical order
		sort.Slice(q, func(i, j int) bool { return q[i] < q[j] })

		// process the first step in the queue
		step := q[0]

		// remove this step from all of its child's incoming degrees
		for i := range edges[step] {
			degrees[edges[step][i]]--
		}

		ans = append(ans, step)
		delete(edges, step)
		delete(degrees, step)
	}
	fmt.Println("Part 1: " + strings.Join(ans, ""))
}

// --- Day 7: The Sum of Its Parts ---

// The instructions specify a series of steps and requirements about which steps
// must be finished before others can begin (your puzzle input). Each step is
// designated by a single letter. For example, suppose you have the following
// instructions:

// Step C must be finished before step A can begin.
// Step C must be finished before step F can begin.
// Step A must be finished before step B can begin.
// Step A must be finished before step D can begin.
// Step B must be finished before step E can begin.
// Step D must be finished before step E can begin.
// Step F must be finished before step E can begin.
// Visually, these requirements look like this:

//   -->A--->B--
//  /    \      \
// C      -->D----->E
//  \           /
//   ---->F-----

// Your first goal is to determine the order in which the steps should be
// completed. If more than one step is ready, choose the step which is first
// alphabetically. In this example, the steps would be completed as follows:

// Only C is available, and so it is done first. Next, both A and F are
// available. A is first alphabetically, so it is done next.

// Then, even though F was available earlier, steps B and D are now also
// available, and B is the first alphabetically of the three.

// After that, only D and F are available. E is not available because only some
// of its prerequisites are complete. Therefore, D is completed next.

// F is the only choice, so it is done next.

// Finally, E is completed.

// So, in this example, the correct order is CABDFE.

// In what order should the steps in your instructions be completed?
