package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Day14() {
	// d14_part1()
	d14_part2()
}

func d14_part2(){
	step_max := 40
	polymer, ins_rules, char_map := parseInput14("input")
	// diff := func(a int, b int) int {
	// 	return a-b
	// }
	// past_count := make([]int,len(char_map))
	// new_count := make([]int,0,len(char_map))
	pair_counts := make(map[string]int)
	pair := ""
	for i := range polymer {
		if i < len(polymer)-1{
			pair = polymer[i: i+2]
			pair_counts[pair] ++
		}
	}
	last_char := polymer[len(polymer)-1:]

	for i:=0;i<step_max;i++{
		pair_counts = run_insert(pair_counts, ins_rules)
	}
	counts := count_pairs(pair_counts, last_char, char_map)
	fmt.Println("Day14, Part2:", counts[len(counts)-1]-counts[0])
}

func run_insert(counts map[string]int, rules map[string]string) map[string]int {
	// NN
	// NCN
	// NC CN
	next_count := make(map[string]int)
	first_pair := ""
	second_pair := ""
	inserted := ""
	for i, v := range counts{
		inserted = rules[i]
		first_pair = i[:1] + inserted
		second_pair = inserted + i[1:]
		if k, ok := next_count[first_pair]; ok{
			next_count[first_pair] = v + k
		} else {
			next_count[first_pair] = v
		}
		if k, ok := next_count[second_pair]; ok{
			next_count[second_pair] = v + k
		} else {
			next_count[second_pair] = v
		}
	}

	return next_count
}

func d14_part1() {
	step_max := 10
	polymer, ins_rules, char_map := parseInput14("input")
	for i:=0;i<step_max;i++{
		polymer = next_step(polymer, ins_rules)
	}
	counts := count_poly(polymer, char_map)
	// fmt.Println(counts)
	fmt.Println("Day14, Part1:",counts[len(counts)-1]-counts[0])
}

func count_pairs(pairs map[string]int, z string, char_map map[rune]int) []int {
	counts := make([]int,len(char_map))

	for k, v := range pairs{
		char_map[rune(k[0])] += v
	}
	char_map[rune(z[0])] ++

	i:=0
	for _,v := range char_map {
		counts[i] = v
		i++
	}
	sort.Ints(counts)
	return counts
}

func count_poly(p string, char_map map[rune]int) []int {
	for _,v := range p {
		char_map[v] ++
	}
	counts := make([]int,len(char_map))
	i:=0
	for _,v := range char_map {
		counts[i] = v
		i++
	}
	sort.Ints(counts)
	return counts
}

func parseInput14(f string) (string, map[string]string, map[rune]int) {
	file, err := os.Open("./data/day14_"+ f +".txt")
	if err != nil {
		fmt.Println(err)
		return "", nil, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	polymer := scanner.Text()
	scanner.Scan()

	ins_rules := make(map[string]string)
	char_map := make(map[rune]int)
	for scanner.Scan() {
		next := strings.Fields(scanner.Text())  // e.g. 00 -> N
		ins_rules[next[0]] = next[2]
		char_map[rune(next[2][0])] = 0
	}
	return polymer, ins_rules, char_map
}

func next_step_linear(p string, ins map[string]string) string{
	var buffer bytes.Buffer
	// final := (len(p) * 2) - 1
	new_poly := make([]string,len(p)-1)

	for i:=1;i<len(p);i++{
		buffer.WriteRune(rune(p[i-1]))
		buffer.WriteRune(rune(ins[p[i-1:i+1]][0]))
		new_poly[i-1] = ins[p[i-1:i+1]]
	}
	buffer.WriteRune(rune(p[len(p)-1]))

	return buffer.String()
}

func next_step(p string, ins map[string]string) string{
	return merge_inserts(p, get_inserts(p, ins))
}

func get_inserts(p string, ins map[string]string) []string {
	new_poly := make([]string,len(p)-1)

	for i:=1;i<len(p);i++{
		new_poly[i-1] = ins[p[i-1:i+1]]
	}
	return new_poly
}

func merge_inserts(p string, ins []string) string {
	merged := ""
	ins_p := append(ins, "")
	for i,v := range p{
		merged += string(v) + ins_p[i]
	}
	return merged
}
