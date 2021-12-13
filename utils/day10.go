package utils

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	slice "github.com/psampaz/slice"
)

func Day10() {
	// d10_part1()
	d10_part2()
}

func d10_part2(){
	file, err := os.Open("./data/day10_input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0, 10)

	for scanner.Scan() {
		new_line := scanner.Text()
		lines = append(lines, new_line)
	}

	// stack with all opening characters
	// value with expected closing
	// value with current opener
	// read each value:
	// 	if opener, add old current to stack, update opener update expected
	// 	if closer and valid, pop from stack and assign to opener, update expected
	stack := make([]rune, 0, len(lines[0]))
	invalids := make([]rune,0)
	remove_line := false
	totals := make([]int,0)
	scores := map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}

	var cur_open rune = '_'
	i := 0
	for  {
		for _, r := range lines[i] {
			if is_opener(r) {
				stack = append(stack, cur_open)
				cur_open = r
				// exp_close = new_closer(r)
			} else {
				if ! is_valid_closer(cur_open, r) {
					invalids = append(invalids, r)
					remove_line = true
					break
				} else {
					cur_open, stack, _ = slice.PopRune(stack)
					// exp_close = new_closer(cur_open)
				}
			}

		}
		if remove_line{
			lines, _ = slice.DeleteString(lines, i)
			remove_line = false
		} else {
			stack = append(stack, cur_open)
			totals = append(totals, complete_line(stack, scores))
			i ++
		}
		if i == len(lines){
			break
		}
		stack, _ = slice.DeleteRangeRune(stack, 0, len(stack))
		cur_open = '_'
	}

	sort.Ints(totals)
	middle := Get_median(totals)

	fmt.Println("Day10, Part2", middle)
}

func d10_part1() {
	file, err := os.Open("./data/day10_input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0, 10)

	for scanner.Scan() {
		new_line := scanner.Text()
		lines = append(lines, new_line)
	}

	// stack with all opening characters
	// value with expected closing
	// value with current opener
	// read each value:
	// 	if opener, add old current to stack, update opener update expected
	// 	if closer and valid, pop from stack and assign to opener, update expected
	stack := make([]rune, 0, len(lines[0]))
	invalids := make([]rune,0)
	// var exp_close rune = '_'
	var cur_open rune = '_'
	for _, line := range lines {
		for _, r := range line {
			if is_opener(r) {
				stack = append(stack, cur_open)
				cur_open = r
				// exp_close = new_closer(r)
			} else {
				if ! is_valid_closer(cur_open, r) {
					invalids = append(invalids, r)
					break
				} else {
					cur_open, stack, _ = slice.PopRune(stack)
					// exp_close = new_closer(cur_open)
				}
			}

		}
	}

	scores := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	total := 0
	for _, v := range invalids {
		total += scores[v]
	}

	fmt.Println("Day10, Part1:", total)
}

func complete_line(stack []rune, scores map[rune]int) int {
	total := 0
	var next rune
	var err error
	for {
		next, stack, err = slice.PopRune(stack)
		if err != nil || next == '_' { break }
		total = total * 5
		total += scores[new_closer(next)]
	}
	return total
}

func is_opener(o rune) bool{
	switch o {
	case '(':
		return true
	case '[':
		return true
	case '{':
		return true
	case '<':
		return true
	default:
		return false
	}
}

func new_closer(c rune) rune {
	switch c {
	case '(':
		 { return ')'  }
	case '[':
		return ']'
	case '{':
		return '}'
	case '<':
		return '>'
	default:
		return '_'
	}

}

func is_valid_closer(e rune, g rune) bool {
	switch e {
	case '(':
		if g == ')' { return true }
	case '[':
		if g == ']' { return true }
	case '{':
		if g == '}' { return true }
	case '<':
		if g == '>' { return true }
	case '_':
		return true // our initial condition
	default:
		return false
	}
	return false
}
