package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "regexp"
)

func Day8() {
	d8_part1()
	d8_part2()
}

func d8_part2(){
	file, err := os.Open("./data/day8_example.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	before_split := make([][]string, 0)
	after_split := make([][]string, 0)
	// digit_counts := make([]int, 10)
	digit_counts := 0

	for scanner.Scan() {
		new_line := strings.Split(scanner.Text(), "|")
		before_split = append(before_split, strings.Fields(new_line[0]))
		after_split = append(after_split, strings.Fields(new_line[1]))
	}

	// top, up-r, up-l, mid, low-r, low-l, bottom
	// with 1 and 7, we can determine top and (u-r / l-r)
	// top
	// (u-r / l-r)
	// with 4 and 8, we can get (bot / l-l)
	// (bot / l-l)
	// u-l and mid are the parts of 4 that aren't (u-r / l-r)
	// (u-l / mid)
	// check the length 6 items
	// 0 is missing (u-r / mid) - diff with 8 gives us mid
		// mid
		// u-l
	// 9 is missing (l-l / bot) - diff with 8 gives l-l
		// l-l
		// bot
	// 6 is missing (u-r / l-r) - diff with 8 gives u-r
		// u-r
		// l-r
	// no need to check length 5 items
	// 3 is missing l-l and (u-l / mid) - diff with 8 and l-l gives u-l
	// 2 is missing
	for _, output := range after_split {
		for _, v := range output{
			switch len(v) {
			case 2, 3, 4, 7:
				digit_counts += 1
			}
		}
	}
	fmt.Println("Day8, Part2")
}

func d8_part1() {
	file, err := os.Open("./data/day8_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	before_split := make([][]string, 0)
	after_split := make([][]string, 0)
	// digit_counts := make([]int, 10)
	digit_counts := 0

	for scanner.Scan() {
		new_line := strings.Split(scanner.Text(), "|")
		before_split = append(before_split, strings.Fields(new_line[0]))
		after_split = append(after_split, strings.Fields(new_line[1]))
	}

	for _, output := range after_split {
		for _, v := range output{
			switch len(v) {
			case 2, 3, 4, 7:
				digit_counts += 1
			}
		}
	}
	fmt.Println("Day8, Part1:", digit_counts)
}
