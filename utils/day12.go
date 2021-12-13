package utils

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Day12() {
	d12_part1()
	d12_part2()
}

func d12_part2(){
	file, err := os.Open("./data/day12_input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	caves_map := make(map[string][]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		new_row := strings.Split(scanner.Text(), "-")

		if _, ok := caves_map[new_row[0]]; ! ok {
			caves_map[new_row[0]] = make([]string,0)
		}
		if _, ok := caves_map[new_row[1]]; ! ok {
			caves_map[new_row[1]] = make([]string,0)
		}

		caves_map[new_row[0]] = add_cave(caves_map[new_row[0]], new_row[1])
		caves_map[new_row[1]] = add_cave(caves_map[new_row[1]], new_row[0])

	}

	routes := 0
	routes = walk_tree_lower("start", make([]string,0), caves_map["start"], caves_map)
	fmt.Println("Day12, Part2:", routes)
}


func d12_part1() {
	file, err := os.Open("./data/day12_input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	caves_map := make(map[string][]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		new_row := strings.Split(scanner.Text(), "-")

		if _, ok := caves_map[new_row[0]]; ! ok {
			caves_map[new_row[0]] = make([]string,0)
		}
		if _, ok := caves_map[new_row[1]]; ! ok {
			caves_map[new_row[1]] = make([]string,0)
		}

		caves_map[new_row[0]] = add_cave(caves_map[new_row[0]], new_row[1])
		caves_map[new_row[1]] = add_cave(caves_map[new_row[1]], new_row[0])

	}

	routes := walk_tree("start", make([]string,0), caves_map["start"], caves_map)
	fmt.Println("Day12, Part1:", routes)
}

func walk_tree_lower(cur string, wib []string, next []string, m map[string][]string) int{
	// fmt.Println(cur,wib,next)
	total := 0
	if cur == "end"{
		// fmt.Println(append(wib,cur))
		return 1
	}
	excl_up := func(s string) bool{
		return (s == strings.ToLower(s))
	}
	no_uppr := FilterString(append(wib, cur), excl_up)
	walk := make([]string, 0)
	if ContainsDupString(no_uppr){
		keep := func(s string) bool{
			return !ContainsString(
				no_uppr, s)
		}
		walk = FilterString(next, keep)
	} else {
		keep := func (s string) bool{
			return s != "start"
		}
		walk = FilterString(next, keep)
	}
	// fmt.Println("    ",no_uppr)
	// fmt.Println("    ",walk)
	for _,v := range walk {
		total += walk_tree_lower(v, append(wib, cur), m[v], m)
	}
	// fmt.Println("returning from",cur)
	return total
}

func walk_tree(cur string, wib []string, next []string, m map[string][]string) int{
	// fmt.Println(cur,wib,next)
	total := 0
	if cur == "end"{
		// fmt.Println(append(wib,cur))
		return 1
	}
	excl_up := func(s string) bool{
		return (s == strings.ToLower(s))
	}
	no_uppr := FilterString(wib, excl_up)
	// fmt.Println("    ",no_uppr)
	keep := func(s string) bool{
		return !ContainsString(
			no_uppr, s)
	}
	walk := FilterString(next, keep)
	// fmt.Println("    ",walk)
	for _,v := range walk {
		total += walk_tree(v, append(wib, cur), m[v], m)
	}
	// fmt.Println("returning from",cur)
	return total
}

func ContainsDupString(x []string) bool {
	sorted := x
	sort.Strings(sorted)
	for i := 1; i<len(sorted); i++{
		if sorted[i] == sorted[i-1]{
			return true
		}
	}
	return false
}

func ContainsString(x []string, s string) bool {
	for _,v := range x {
		if v == s {
			return true
		}
	}
	return false
}

func FilterString(x []string, f func(s string) bool) []string {
	filtered := make([]string,0,len(x))
	for _,v := range x {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func add_cave(caves []string, new string) []string {
	exists := false
	for _, v := range caves {
		if v == new {
			exists = true
		}
	}
	if !exists{
		return append(caves, new)
	} else {
		return caves
	}
}

