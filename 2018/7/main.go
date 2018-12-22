package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bwiggs/advent-of-code/2018/lib"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	edges := make(map[string][]string)
	degrees := make(map[string]int)
	for _, l := range lib.ReadLines("test-input.txt") {
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

	q := []string{}
	ans := []string{}
	for {
		// queue up all ready tasks
		for step := range edges {
			if degrees[step] == 0 {
				q = append(q, step)
			}
		}

		// if there's not tasks, end
		if len(q) == 0 {
			break
		}

		sort.Slice(q, func(i, j int) bool { return q[i] < q[j] })

		// process each step in the queue
		for _, step := range q {
			ans = append(ans, step)
			// for k, v := range edges[step] {
			// }
		}
		break
	}
	fmt.Println("Part 1: " + strings.Join(ans, ""))
}

func mainOld() {
	graph := buildGraph()
	graph.Print()
}

func buildGraph() *Node {
	var root *Node
	list := []string{}
	for _, l := range lib.ReadLines("input.txt") {
		fmt.Printf("\n----- %s -----\n", l)
		words := strings.Split(l, " ")

		parentName := words[1]
		childName := words[7]

		list = append(list, parentName+childName)

		if root == nil {
			fmt.Printf("root: %s\n", parentName)
			root = NewNode(parentName)
		}

		var depNode *Node
		if depNode = root.Find(childName); depNode == nil {
			fmt.Printf("creating: %s\n", childName)
			depNode = NewNode(childName)
		}

		fmt.Printf("inserting: %s -> %s\n", depNode.Name, parentName)
		if root == nil {
			root = NewNode(parentName)
		}
		root.Insert(parentName, depNode)
	}

	sort.Slice(list, func(i, j int) bool { return list[i] < list[j] })
	fmt.Println(list)

	spew.Dump(root)

	return root
}

// Node type
type Node struct {
	Name     string
	Children []*Node
	Prereqs  int
	Visited  bool
}

// NewNode returns a new Node type
func NewNode(name string) *Node {
	return &Node{name, []*Node{}, 0, false}
}

// Find node
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

// Print node tree
func (n *Node) Print() {
	fmt.Print(n.Name)

	// remove 1 dependnency from all children
	for _, c := range n.Children {
		c.Prereqs--
		if !c.Visited && c.Prereqs == 0 {
			c.Print()
			// nc++
		}
	}

	// execute any children that are ready
	// nc := 0
	// for i := 0; nc != len(n.Children); i++ {
	// 	if !n.Children[i].Visited && n.Children[i].Prereqs == 0 {
	// 		n.Children[i].Print()
	// 		nc++
	// 	}
	// 	if i == len(n.Children)-1 {
	// 		i = -1
	// 	}
	// }
	// n.Visited = true
}

// Insert node
func (n *Node) Insert(nodeName string, child *Node) (inserted bool) {
	// fmt.Printf("visiting: %s\n", n.Name)
	if n.Name == nodeName {
		n.Children = append(n.Children, child)
		sort.Slice(n.Children, func(i, j int) bool { return n.Children[i].Name < n.Children[j].Name })
		child.Prereqs++
		return true
	}

	for _, c := range n.Children {
		if c.Insert(nodeName, child) {
			return true
		}
	}

	return false
}

func (n *Node) String() string {
	s := fmt.Sprintf("%s: ", n.Name)
	for _, c := range n.Children {
		s += c.Name
	}
	return s
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
