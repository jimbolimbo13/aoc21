package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	maxY int
)

func Day17() {
	// d17_part1()
	d17_part2()
}

func d17_part2(){
	xLimits, yLimits := parseInput17("input")
	// fmt.Println(xLimits, yLimits)

	target_range := make([][]int,2)
	target_range[0] = xLimits
	target_range[1] = yLimits

	yRange := 2000
	guesses := make([][]int,yRange)
	for i := range guesses {
		guesses[i] = make([]int, target_range[0][1]+1)
	}

	maxY = 0
	target_count := 0
	for i, v := range guesses {
		for j := range v {
			if isReachTarget(j, i-(yRange/2), target_range){
				guesses[i][j] = 1
				target_count ++
			}
		}
	}

	// for _,v  := range guesses {
	// 	for _,k := range v {
	// 		fmt.Print(k," ")
	// 	}
	// 	fmt.Println()
	// }

	// fmt.Println(isReachTarget(6, 9, target_range))
	fmt.Println("Day17, Part2:", target_count)
}

func d17_part1() {
	xLimits, yLimits := parseInput17("input")
	// fmt.Println(xLimits, yLimits)

	target_range := make([][]int,2)
	target_range[0] = xLimits
	target_range[1] = yLimits

	guesses := make([][]int,1000)
	for i := range guesses {
		guesses[i] = make([]int, 1000)
	}

	maxY = 0
	for i, v := range guesses {
		for j := range v {
			if isReachTarget(i, j, target_range){
				guesses[i][j] = 1
			}
		}
	}

	// for _,v  := range guesses {
	// 	for _,k := range v {
	// 		fmt.Print(k," ")
	// 	}
	// 	fmt.Println()
	// }

	// fmt.Println(isReachTarget(6, 9, target_range))
	fmt.Println("Day17, Part1:", maxY)
}

func isReachTarget( xVel int, yVel int, target [][]int ) bool {
	localMaxY := 0
	xPos := 0
	yPos := 0
	next := make([]int, 4)
	for {
		if yPos > localMaxY { localMaxY = yPos }
		if xPos >= target[0][0] && xPos <= target[0][1] && yPos >= target[1][0] && yPos <= target[1][1] {
			if localMaxY > maxY { maxY = localMaxY }
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
