package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day12() {
	d12_part1()
	d12_part2()
}

func d12_part2(){
	fmt.Println("Day12, Part2:")
}

type cave struct {
	name string
	next []cave // change this to a map to emulate a set
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
	first_cave := cave{"start", make([]cave,0)}
	add_nodes(&first_cave, caves_map)
	for i, _ := range first_cave.next{
		add_nodes(&(first_cave.next[i]), caves_map)
	}

	routes := walk_tree("start", make([]string,0), caves_map["start"], caves_map)
	// walk_tree(&caves_map, "start", "_", "")
	fmt.Println("Day12, Part1:", routes)
}

func walk_tree(cur string, wib []string, next []string, m map[string][]string) int{
	// fmt.Println(cur,wib,next)
	total := 0
	if cur == "end"{
		// fmt.Println(append(wib,cur))
		return 1
	}
	excl_up := func(s string) bool{
		return s == strings.ToLower(s)
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

func add_nodes(c *cave, m map[string][]string) {
	for _, v := range m[c.name]{
		(*c).next = add_node((*c).next, v)
	}
}

func add_node(nodes []cave, new string) []cave {
	exists := false
	for _, v := range nodes {
		if v.name == new {
			exists = true
		}
	}
	if !exists{
		return append(nodes, cave{new,make([]cave,0)})
	} else {
		return nodes
	}
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

// old stuff

// func walk_tree_o(l *map[string]cave, n string, called string, tree string){
// 	fmt.Println((*l)[n].name, (*l)[n])
// 	fmt.Println(tree,n)
// 	if n != strings.ToUpper(n) && n != "start"{
// 		(*l)[called].next[n] = false
// 	}
// 	if n == "end" { return }
// 	if i, ok := (*l)[n]; ok {
// 		for k, t := range i.next  {
// 			fmt.Println("in",k)
// 			if k != "start" && t {
// 				walk_tree(l, k, n, (tree+","+n))
// 			}
// 		}
// 	}
// }


// func print_node(n cave, pfx string){
// 	fmt.Println(pfx,n.name)
// 	if n.name != "start" {
// 		for i,_ := range n.next {
// 			if  i != "end" {
// 				new_pfx := pfx + "| "
// 				// print_node(v, new_pfx)
// 				fmt.Println(new_pfx)
// 			}
// 		}

// 	}
// }

// func print_node_r(n cave, pfx string){
// 	fmt.Printf("%s,%s\n",n.name,pfx)
// 	if n.name != "start" {
// 		for i,_ := range n.next {
// 			if  i != "end" {
// 				new_pfx := n.name + "," + pfx
// 				//print_node_r(v, new_pfx)
// 				fmt.Println(new_pfx)
// 			}
// 		}

// 	}

// }



// func link_caves(c1 *cave, c2 *cave) {
// 	if ! is_linked(*c1, *c2) {
// 		(*c1).next[(*c2).name] = true
// 	}
// 	if ! is_linked(*c2, *c1) {
// 		(*c2).next[(*c1).name] = true
// 	}
// 	if (*c1).name == strings.ToUpper((*c1).name){
// 		(*c1).next[(*c2).name] = true
// 	}
// 	if (*c2).name == strings.ToUpper((*c2).name){
// 		(*c2).next[(*c1).name] = true
// 	}
// }

// func is_linked(c1 cave, c2 cave) bool {
// 	linked := false
// 	for i := range c1.next {
// 		if i == c2.name {
// 			linked = true
// 		}
// 	}
// 	return linked
// }
