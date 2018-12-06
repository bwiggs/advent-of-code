package main

import (
	"fmt"
	"io/ioutil"
)

// --- Day 5: Alchemical Reduction ---

// first time i wrote this was with the linear approach using a for loop. It
// modified the list in line. The stack approach is much cleaner, but uses additional
// space to for stack storage.

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	polymer := append(dat[:0:0], dat...)
	fmt.Printf("Part 1 w/ stack: %d\n", reactStack(polymer, 0))

	for i := 65; i < 91; i++ {
		polymer = append(dat[:0:0], dat...)
		fmt.Printf("Part 2 w/ stack: %d ignore:%s\n", reactStack(polymer, byte(i)), string(i))
	}
}

func reactStack(polymer []byte, ignore byte) int {
	stack := []byte{}
	for _, b := range polymer {

		if b == ignore || b == ignore+32 {
			continue
		}

		if len(stack) == 0 {
			stack = append(stack, b)
			continue
		}

		top := len(stack) - 1
		if stack[top]+32 == b || stack[top]-32 == b {
			// pop the stack
			stack = stack[:top]
		} else {
			stack = append(stack, b)
		}
	}
	return len(stack)
}

// func main_polymer_inline() {
// 	dat, _ := ioutil.ReadFile("input.txt")
// 	polymer := append(dat[:0:0], dat...)
// 	fmt.Printf("Part 1: %d\n", react(polymer))

// 	polymer = append(dat[:0:0], dat...)
// 	var p string
// 	pl := 9999999
// 	for i := 65; i < 91; i++ {
// 		ln := part2(polymer, byte(i))
// 		if ln < pl {

// 			p = string(i)
// 			pl = ln
// 		}
// 	}
// 	fmt.Printf("Part 2: %s %d\n", p, pl)
// }

// modifies the polymer inline, some nasty iteration variable management
// func react(polymer []byte) int {
// 	for i := 0; i < len(polymer)-1; i++ {
// 		if polymer[i]+32 == polymer[i+1] || polymer[i]-32 == polymer[i+1] {
// 			end := lib.Min(2, len(polymer)-1-i)
// 			polymer = append(polymer[:i], polymer[i+end:]...)
// 			i = lib.Max(-1, i-2)
// 		}
// 	}
// 	return len(polymer)
// }

// func part2(polymer []byte, b byte) int {
// 	polymer = bytes.Replace(polymer, []byte{b}, []byte{}, -1)
// 	polymer = bytes.Replace(polymer, []byte{b + 32}, []byte{}, -1)
// 	return reactStack(polymer)
// }

// --- Day 5: Alchemical Reduction ---

// You've managed to sneak in to the prototype suit manufacturing lab. The Elves
// are making decent progress, but are still struggling with the suit's size
// reduction capabilities.

// While the very latest in 1518 alchemical technology might have solved their
// problem eventually, you can do better. You scan the chemical composition of
// the suit's material and discover that it is formed by extremely long polymers
// (one of which is available as your puzzle input).

// The polymer is formed by smaller units which, when triggered, react with each
// other such that two adjacent units of the same type and opposite polarity are
// destroyed. Units' types are represented by letters; units' polarity is
// represented by capitalization. For instance, r and R are units with the same
// type but opposite polarity, whereas r and s are entirely different types and
// do not react.

// For example:

// In aA, a and A react, leaving nothing behind.
// In abBA, bB destroys itself, leaving aA. As above, this then destroys itself,
// 	leaving nothing.
// In abAB, no two adjacent units are of the same type, and so nothing happens.
// In aabAAB, even though aa and AA are of the same type, their polarities
// 	match, and so nothing happens.

// Now, consider a larger example, dabAcCaCBAcCcaDA:

// dabAcCaCBAcCcaDA  The first 'cC' is removed.
// dabAaCBAcCcaDA    This creates 'Aa', which is removed.
// dabCBAcCcaDA      Either 'cC' or 'Cc' are removed (the result is the same).
// dabCBAcaDA        No further actions can be taken.

// After all possible reactions, the resulting polymer contains 10 units.

// How many units remain after fully reacting the polymer you scanned? (Note: in
// this puzzle and others, the input is large; if you copy/paste your input,
// make sure you get the whole thing.)
