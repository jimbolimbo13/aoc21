package utils

import (
	"fmt"
	"strconv"
	"strings"
	"os"
	"bufio"
)

func Day2() {
	type position struct {
		distance int
		depth int
		aim int
	}
	sub := position{0, 0, 0}

	// data := []string{"forward 5","down 5","forward 8","up 3","down 8","forward 2"}

	file, err := os.Open("./data/day2_data.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
	// for _, v := range data {
	//  	split := strings.Split(v, " ")
		direction := split[0]
		amount, _:= strconv.Atoi(split[1])

		switch direction {
		case "forward":
			sub.distance += amount
			sub.depth += sub.aim * amount
		case "down":
			sub.aim += amount
		case "up":
			sub.aim -= amount
		}

	}

	fmt.Println("Day2, Part2", sub.distance * sub.depth)
}

// func move(direction string, amount int, sub *position) {

// 	switch direction {
// 	case "forward":
// 		*sub.distance += amount
// 	case "down":
// 		*sub.depth += amount
// 	case "up":
// 		*sub.depth -= amount
// 	}
// }
