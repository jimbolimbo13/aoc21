package utils

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"sync"
	"strconv"
)

var (
	growth_period = 80
)

func Day6() {
	// d6_part1()
	d6_part2()
}

func d6_part2(){
	growth_period = 256
	d6_part1()
	fmt.Println("Day6, Part2")
}

func d6_part1() {
	file, err := os.Open("./data/day6_example.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var growing sync.WaitGroup
	var birth sync.Mutex
	scanner := bufio.NewScanner(file)
	angler_school := make([]int,0,1000)
	for scanner.Scan() {
		og_fish := strings.Split(scanner.Text(), ",")
		for _, v := range og_fish {
			next_fish, _ := strconv.Atoi(v)
			angler_school = append(angler_school, next_fish)
		}
	}
	orig := len(angler_school)
	day := 0
	for i := 0; i < orig; i ++{
		growing.Add(1)
		go angler_fish(&angler_school, i, day, &growing, &birth)
	}

	growing.Wait()
	// fmt.Println(angler_school)
	fmt.Println("Day6, Part1:", len(angler_school))
}

// func wait_for_growth(growing *sync.WaitGroup)

func angler_fish(school *[]int, pos int, start int, growing *sync.WaitGroup, birth *sync.Mutex){
	new_pos := 0
	for i := start ; i < growth_period; i ++ {
		// if pos == 0 {
		// 	fmt.Println("Day:",i)
		// }
		if (*school)[pos] == 0 {
			birth.Lock()
			*school = append(*school, 8)
			new_pos = len(*school)
			birth.Unlock()
			growing.Add(1)
			go angler_fish(school, new_pos - 1, i + 1, growing, birth)
			// if pos == 0 {
			// 	fmt.Println("new_pos",(*school)[new_pos - 1])
			// }
			(*school)[pos] = 6
		} else {
			(*school)[pos] -= 1
		}
	}
	growing.Done()
}
