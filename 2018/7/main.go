package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bwiggs/advent-of-code/2018/lib"
)

type Node struct {
	Name     string
	Children []*Node
	Prereqs  int
}

func main() {
	graph := getNodes()
	graph.Print()
}

func getNodes() *Node {
	var root *Node
	for _, l := range lib.ReadLines("test-input.txt") {
		words := strings.Split(l, " ")

		nodeName := words[1]
		childName := words[7]

		if root == nil {
			// fmt.Println("created root node")
			root = &Node{nodeName, []*Node{&Node{childName, []*Node{}}}}
			continue
		}

		// fmt.Println("new node read!")

		var depNode *Node
		if depNode = root.Find(childName); depNode == nil {
			// fmt.Printf("creating new node: %s\n", childName)
			depNode = NewNode(childName)
		}

		// fmt.Printf("inserting: %s -> %s\n", nodeName, depNode.Name)
		root.Insert(nodeName, depNode)
	}
	return root
}

func (n *Node) Find(nodeName string) *Node {
	if n.Name == nodeName {
		return n
	}

	for _, c := range n.Children {
		if nn := c.Find(nodeName); nn != nil {
			return nn
		}
	}

	return nil
}

func (n *Node) Print() {
	fmt.Print(n.Name)
	for _, c := range n.Children {
		c.Prereqs++
	}
}

func (n *Node) Insert(nodeName string, child *Node) (inserted bool) {
	// fmt.Printf("visiting: %s\n", n.Name)
	if n.Name == nodeName {
		n.Children = append(n.Children, child)
		sort.Slice(n.Children, func(i, j int) bool { return n.Children[i].Name < n.Children[j].Name })
		return true
	}

	for _, c := range n.Children {
		if c.Insert(nodeName, child) {
			return true
		}
	}

	return false
}

func NewNode(name string) *Node {
	return &Node{name, []*Node{}}
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
