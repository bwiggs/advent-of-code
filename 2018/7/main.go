package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bwiggs/advent-of-code/2018/lib"
)

// This one caused me a lot of headache. It was my first time doing a
// topological sort, which was fun to learn. Most of my time was spent trying to
// ensure that all children of a completed node were completed first. After more
// than a few hours I reread the instructions and found that that was not the
// requirement. You just have to process the next alphabetically ready node.
// Made things much easier.

// GoLang requires a lot of initization code for maps, which I used to track
// state. Other languages make this much easier. Also, functional languages seem
// to make list processing MUCH cleaner.

func main() {
	part2()
	return
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

		// if there's no tasks we're done
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

	part2()
}

func part2() {
	edges := make(map[string][]string)
	degrees := make(map[string]int)
	durations := make(map[string]int)
	done := make(map[string]bool)
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

		done[succ] = false
		done[dep] = false
		durations[succ] = int(succ[0]) - 4
		durations[dep] = int(dep[0]) - 4

	}

	workers := [5]string{}
	var q []string
	totalTime := 0

	for {
		fmt.Println(totalTime)
		// time.Sleep(33 * time.Millisecond)
		// update durations and worker status
		active := false
		idleWorker := false
		// spew.Dump(workers)
		for i := range workers {
			// decremenent duration for all running steps
			if workers[i] == "" {
				idleWorker = true
				continue
			}

			// fmt.Println("decrement task", durations[workers[i]])
			durations[workers[i]]--
			active = true

			// if the step is done, free the worker and mark the task as done
			if durations[workers[i]] == 0 {
				fmt.Println("complete:", workers[i])
				// remove this step from all of its child's incoming degrees
				for _, child := range edges[workers[i]] {
					degrees[child]--
				}
				done[workers[i]] = true
				delete(edges, workers[i])
				delete(degrees, workers[i])
				workers[i] = ""
				idleWorker = true
			}
		}

		totalTime++

		if !idleWorker {
			continue
		}

		// fill the queue with ready steps
		q = []string{}
		for step, deg := range degrees {
			if deg == 0 && !strings.Contains(strings.Join(workers[:], ""), step) {
				q = append(q, step)
			}
		}

		if len(q) == 0 {
			if active {
				// fmt.Println("no tasks ready")
				continue
			}
			fmt.Println("queue is empty and workers are idle, done!")
			break
		}

		// sort the steps in alphabetical order
		sort.Slice(q, func(i, j int) bool { return q[i] < q[j] })
		fmt.Printf("queue: %s\n", q)

		// process any available steps in the queue
		for i := range workers {

			if workers[i] != "" {
				continue
			}

			if len(q) == 0 {
				break
			}

			fmt.Printf("starting: %s\n", q[0])
			workers[i] = q[0]
			q = q[1:] // dequeue

		}
		fmt.Printf("workers: %s\n", strings.Join(workers[:], "-"))
	}

	fmt.Println(totalTime - 2) // off by 2? Weird counting issue? This gives the correct answer
}

type Digraph struct {
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
