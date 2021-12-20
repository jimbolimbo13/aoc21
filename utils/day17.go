package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day17() {
	d17_part1()
	d17_part2()
}

func d17_part2(){
	parseInput17("example")
	fmt.Println("Day17, Part2:")
}

func d17_part1() {
	xLimits, yLimits := parseInput17("example")
	fmt.Println(xLimits, yLimits)

	target_range := make([][]int,2)
	target_range[0] = xLimits
	target_range[1] = yLimits

	fmt.Println(isReachTarget(6, 9, target_range))
	fmt.Println("Day17, Part1:")
}

func isReachTarget( xVel int, yVel int, target [][]int ) bool {
	xPos := 0
	yPos := 0
	next := make([]int, 4)
	fmt.Println("is", target)
	for {
		if xPos >= target[0][0] && xPos <= target[0][1] && yPos >= target[1][0] && yPos <= target[1][1] {
			return true
		} else if xPos > target[0][1] || yPos < target[1][0] {
			return false
		}
		next = nextStep(xPos, yPos, xVel, yVel)
		xPos = next[0]
		yPos = next[1]
		xVel = next[2]
		yVel = next[3]
	}
}

func nextStep (xPos int, yPos int, xVel int, yVel int) []int {
	xPos += xVel
	if xVel > 0 { xVel -- }
	yPos += yVel
	yVel --
	fmt.Println("nextStep", xPos, yPos)
	return []int{xPos, yPos, xVel, yVel}
}

func parseInput17(f string) ([]int, []int) {
	file, err := os.Open("./data/day17_"+ f +".txt")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	xCoords := make([]string,0)
	yCoords := make([]string,0)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), ":")
		xyText := strings.Split(text[1], ",")
		xText := strings.Split(xyText[0], "=")
		xCoords = strings.Split(xText[1], "..")

		yText := strings.Split(xyText[1], "=")
		yCoords = strings.Split(yText[1], "..")
	}
	xLimits := make([]int, 2)
	xLimits[0], _ = strconv.Atoi(xCoords[0])
	xLimits[1], _ = strconv.Atoi(xCoords[1])

	yLimits := make([]int, 2)
	yLimits[0], _ = strconv.Atoi(yCoords[0])
	yLimits[1], _ = strconv.Atoi(yCoords[1])

	return xLimits, yLimits
}
