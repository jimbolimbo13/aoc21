package utils

import (
	"fmt"
	"bufio"
	"os"
)

func DayXX() {
	dXX_part1()
	dXX_part2()
}

func dXX_part2(){
	fmt.Println("DayXX, Part2")
}

func dXX_part1() {
	file, err := os.Open("./data/day4_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// process data
	}
	fmt.Println("DayXX, Part1:")
}
