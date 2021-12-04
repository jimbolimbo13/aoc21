package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	results = make([]int, 0)
	playing sync.WaitGroup
)

func Day4() {
	d4_part1()
	d4_part2()
}

func d4_part2(){

	file, err := os.Open("./data/day4_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	bingo_draws := strings.Split(scanner.Text(), ",")

	scanner.Scan()

	boards := make([][]int, 0)
	board := make([]int, 0)
	b := 0

	for scanner.Scan() {
		split := strings.Fields(scanner.Text())
		for _, v := range split{
			next, _ := strconv.Atoi(v)
			board = append(board, next)
		}
		if len(split) != 0 {
			b ++
		}
		if b == 5 {
			boards = append(boards, board)
			shadow_board := make([]int, len(board))
			boards = append(boards, shadow_board)
			board = make([]int, 0)
			b = 0
		}

	}

	results = make([]int, len(boards)/2)
	playing.Add(len(boards)/2)
	for i := 0; i < len(boards); i += 2{
		pos := (i*1/2)
		go play_board(boards[i], boards[i+1], bingo_draws, pos)

	}
	playing.Wait()
	slowest, last_draw := get_slowest(len(bingo_draws))
	board_total := calculate_winner(boards[slowest*2], boards[1+slowest*2])

	ld_int, _ := strconv.Atoi(bingo_draws[last_draw])
	fmt.Println("Day4, Part2:", board_total * ld_int)
}

func d4_part1() {
	file, err := os.Open("./data/day4_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	bingo_draws := strings.Split(scanner.Text(), ",")

	scanner.Scan()

	boards := make([][]int, 0)
	board := make([]int, 0)
	b := 0

	for scanner.Scan() {
		split := strings.Fields(scanner.Text())
		for _, v := range split{
			next, _ := strconv.Atoi(v)
			board = append(board, next)
		}
		if len(split) != 0 {
			b ++
		}
		if b == 5 {
			boards = append(boards, board)
			shadow_board := make([]int, len(board))
			boards = append(boards, shadow_board)
			board = make([]int, 0)
			b = 0
		}

	}

	results = make([]int, len(boards)/2)
	playing.Add(len(boards)/2)
	for i := 0; i < len(boards); i += 2{
		pos := (i*1/2)
		go play_board(boards[i], boards[i+1], bingo_draws, pos)

	}
	playing.Wait()
	fastest, last_draw := get_fastest(len(bingo_draws))
	board_total := calculate_winner(boards[fastest*2], boards[1+fastest*2])

	ld_int, _ := strconv.Atoi(bingo_draws[last_draw])
	fmt.Println("Day4, Part1:", board_total * ld_int)
}

func calculate_winner(board []int, shadow []int) int {
	board_total := 0
	for i, v := range shadow {
		board_total += (1^v)*board[i]
	}
	return board_total
}

func get_slowest(drawing_count int) (int, int) {
	slowest := 0
	position := -1
	for i, v := range results {
		if v > slowest {
			slowest = v
			position = i
		}
	}
	return position, slowest
}

func get_fastest(drawing_count int) (int, int) {
	fastest := drawing_count
	position := -1
	for i, v := range results {
		if v < fastest {
			fastest = v
			position = i
		}
	}
	return position, fastest
}

func play_board(board []int, shadow []int, drawings []string, position int) int {
	for i, v := range drawings {
		draw, _ := strconv.Atoi(v)
		match := check_draw(board, draw)
		if match != -1 {
			shadow[match] = 1
			if check_bingo(shadow){
				results[position] = i
				playing.Done()
				return i
			}
		}
	}
	return -1
}

func check_bingo(shadow []int) bool {
	ret := false

	for i := 0; i < 5; i++{
		column := []int{shadow[(i%5)], shadow[(i%5)+5], shadow[(i%5)+10], shadow[(i%5)+15], shadow[(i%5)+20]}
		row := []int{shadow[(i*5)], shadow[(i*5)+1], shadow[(i*5)+2], shadow[(i*5)+3], shadow[(i*5)+4]}
		if (sum(row) == 5) || (sum(column) == 5){
			return true
		}
	}
	return ret
}

func check_draw(board []int, draw int) int {
	for i, v := range board {
		if v == draw {
			return i
		}
	}
	return -1
}
