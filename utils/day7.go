package utils

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"strconv"
	"math"
)

func Day7() {
	d7_part1()
	d7_part2()
}

func d7_part2(){
	file, err := os.Open("./data/day7_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	crabs := make([]int, 0)
	for scanner.Scan() {
		crabs_import := strings.Split(scanner.Text(), ",")
		for _, v := range crabs_import {
			next_crab, _ := strconv.Atoi(v)
			crabs = append(crabs, next_crab)
		}
	}

	sort.Ints(crabs)

	gather_point := Get_median(crabs)
	first_cost := get_fleet_cost(crabs, gather_point)
	lower_cost := get_fleet_cost(crabs, gather_point-1)
	higher_cost := get_fleet_cost(crabs, gather_point+1)
	best_cost := 0

	dir := 0
	if higher_cost < first_cost {
		dir = 1
		best_cost = higher_cost
	} else if lower_cost < first_cost {
		dir = -1
		best_cost = lower_cost
	}

	next_cost := 0
	if dir != 0 {
		i := gather_point + dir*2
		for {
			next_cost = get_fleet_cost(crabs, i)
			if next_cost < best_cost{
				i += dir
				best_cost = next_cost
			} else {
				break
			}

		}
	}

	fmt.Println("Day7, Part2", best_cost)
}

func d7_part1() {
	file, err := os.Open("./data/day7_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	crabs := make([]int, 0)
	for scanner.Scan() {
		crabs_import := strings.Split(scanner.Text(), ",")
		for _, v := range crabs_import {
			next_crab, _ := strconv.Atoi(v)
			crabs = append(crabs, next_crab)
		}
	}

	sort.Ints(crabs)
	// crabs_grouped := make([]int, crabs[len(crabs)-1]+1)

	gather_point := Get_median(crabs)
	fuel_cost := 0
	for _,v := range crabs {
		// crabs_grouped[v] += 1
		fuel_cost += int(math.Abs(float64(v-gather_point)))
	}
	fmt.Println("Day7, Part1:", fuel_cost)
}

func get_fleet_cost (fleet []int, gather_point int) int {
	fleet_cost := 0
	for _,v := range fleet {
		fleet_cost += get_cost(int(math.Abs(float64(v-gather_point))))
	}
	return fleet_cost
}

func get_cost(dist int) int {
	cost := 0
	for i := 1; i <= dist ; i++ {
		cost += i
	}
	return cost
}

func Get_median(sl []int) int {
	median := 0
	median = len(sl)/2
	if len(sl) %2 ==1{
		return sl[median]
	} else {
		return (sl[median -1] + sl[median])/2
	}
}
