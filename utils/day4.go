package utils

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"strconv"
)

func Day4() {
	d4_part1()
	////  d4_part2()
}

func d4_part2(){
	fmt.Println("Hello, day4!")
}

func d4_part1() {
	file, err := os.Open("./data/day4_example.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	bingo_draws := strings.Split(scanner.Text(), ",")
	fmt.Println(bingo_draws)

	scanner.Scan()

	boards := make([][]int, 0)
	board := make([]int, 0)

	for scanner.Scan() {
		split := strings.Fields(scanner.Text())
		for _, v := range split{
			next, _ := strconv.Atoi(v)
			board = append(board, next)
		}
		if len(split) == 0 {
			boards = append(boards, board)
			board = make([]int, 0)
		}

	}

	fmt.Println("Hello, day4!")
}
