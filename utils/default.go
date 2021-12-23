package utils

import (
	"bufio"
	"fmt"
	"os"
)

func DayXX() {
	dXX_part1()
	dXX_part2()
}

func dXX_part2(){
	parseInputXX("example")
	fmt.Println("DayXX, Part2:")
}

func dXX_part1() {
	parseInputXX("example")
	fmt.Println("DayXX, Part1:")
}

func parseInputXX(f string) {
	file, err := os.Open("./data/dayXX_"+ f +".txt")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// process data
	}
}
