package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"

	slice "github.com/psampaz/slice"
)

func Day9() {
	d9_part1()
	d9_part2()
}

func d9_part2(){
	file, err := os.Open("./data/day9_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cave_floor := make([][]int, 0)
	basins := make([]int, 0, 100)
	dimensions := make([]int, 2)
	for scanner.Scan() {
		new_row := strings.Split(scanner.Text(), "")
		dimensions[0] += 1
		dimensions[1] = len(new_row)
		next_floor := make([]int, len(new_row))
		for i, v := range new_row {
			next, _ := strconv.Atoi(v)
			next_floor[i] = next
			basins = append(basins, next)
		}
		cave_floor = append(cave_floor, next_floor)

	}

	slopes := make(map[int][]int) // key: lower, value: upper
	has_lower := make(map[int]bool)
	// basin_cnt := 0
	// var prev_slope int = 0

	// lowest_points := make([]int, 0)
	for i, floor := range cave_floor {
		row_len := len(floor)
		for j := range floor {
			here := j + i*row_len
			if basins[here] == 9 { continue }
			if is_valid_hor(here, here+1, dimensions) && basins[here +1] != 9{
				next_slope := slope(basins[here], basins[here + 1] )
				switch next_slope {
				case -1:
					if _, ok := has_lower[here]; !ok{
						add_slope(here+1, here, &slopes)
						has_lower[here] = true
					}
				case 1:
					if _, ok := has_lower[here+1]; !ok{
						add_slope(here, here+1, &slopes)
						has_lower[here+1] = true
					}
				}
			}
			if is_valid_vert(here, here+row_len, dimensions) && basins[here+row_len] !=9 {
				next_slope := slope(basins[here], basins[here + row_len] )
				switch next_slope {
				case -1:
					if _, ok := has_lower[here]; !ok {
						add_slope(here+row_len, here, &slopes)
						has_lower[here] = true
					}
				case 1:
					if _, ok := has_lower[here+row_len]; !ok {
						add_slope(here, here+row_len, &slopes)
						has_lower[here+row_len] = true
					}
				}
			}
			// if j == 3 {break}
		}
	}

	lowest := make([]int,0)
	for lo := range slopes {
		if _, ok := has_lower[lo]; !ok{
			lowest = append(lowest, lo)
		}
	}

	biggest := make([]int, 0)
	for _, v := range lowest {
		total := sum_tree(v, &slopes)
		min, _ := slice.MinInt(biggest)
		if total > min && len(biggest) == 3  {
			new := slice.FilterInt(biggest, func(x int) bool { return x > min })
			biggest = append(new, total)
		} else if len(biggest) < 3 {
			biggest = append(biggest, total)
		}
	}

	// notes for min

	fmt.Println("Day9, Part2:", biggest[0]*biggest[1]*biggest[2])
}

func d9_part1() {
	file, err := os.Open("./data/day9_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cave_floor := make([][]int, 0)
	for scanner.Scan() {
		new_row := strings.Split(scanner.Text(), "")
		next_floor := make([]int, len(new_row))
		for i, v := range new_row {
			next, _ := strconv.Atoi(v)
			next_floor[i] = next
		}
		cave_floor = append(cave_floor, next_floor)

	}

	lowest_points := make([]int, 0)
	for i, floor := range cave_floor {
		for j := range floor {
			if is_lowest(cave_floor, i, j){
				lowest_points = append(lowest_points, cave_floor[i][j] + 1)
			}
		}
	}

	fmt.Println("Day9, Part1:", sum(lowest_points))
}

func sum_tree(head int, slopes *map[int][]int) int {
	total := 1
	if val, ok := (*slopes)[head]; ok{
		for _, v := range val{
			total += sum_tree(v, slopes)
		}
	}
	return total
}

func add_slope(a int, b int, slopes *map[int][]int) {
	if slope, ok := (*slopes)[a]; ok {
		(*slopes)[a] = append(slope, b)
	} else {
		(*slopes)[a] = []int{b}
	}
}

func slope(a int, b int) int {
	if a > b {
		return -1
	} else if b > a {
		return 1
	} else {
		return 0
	}
}

func is_valid_hor(a int, b int, dim []int) bool {
	// 9 - 10
	switch {
		case a < 0:
			return false
		case b < 0:
			return false
		case b % dim[1] == 0:
			return false
		default:
			return true
	}
}

func is_valid_vert(a int, b int, dim []int) bool {
	switch {
		case a < 0:
			return false
		case b < 0:
			return false
		case b >= dim[0]*dim[1]:
			return false
		default:
			return true
	}
}

func is_lowest(cave_floor [][]int, i int, j int) bool {
	// top left corner
	check := cave_floor[i][j]
	if i == 0 && j == 0 {
		if check < cave_floor[i+1][j] &&
			check < cave_floor[i][j+1] {
			return true
		}
	} else if i == 0 && j == (len(cave_floor[i])-1) {
		if check < cave_floor[i+1][j] &&
			check < cave_floor[i][j-1] {
			return true
		}
	} else if i == (len(cave_floor)-1) && j == 0 {
		if check < cave_floor[i-1][j] &&
			check < cave_floor[i][j+1] {
			return true
		}
	} else if i == (len(cave_floor)-1) && j == (len(cave_floor[i])-1) {
		if check < cave_floor[i-1][j] &&
			check < cave_floor[i][j-1]{
				return true
			}
	} else if i == 0 {
		if check < cave_floor[i][j-1] &&
			check < cave_floor[i][j+1] &&
			check < cave_floor[i+1][j] {
			return true
		}
	} else if j == 0 {
		if check < cave_floor[i-1][j] &&
			check < cave_floor[i][j+1] &&
			check < cave_floor[i+1][j] {
			return true
		}
	} else if i >= (len(cave_floor)-1) {
		if check < cave_floor[i-1][j] &&
			check < cave_floor[i][j+1] &&
			check < cave_floor[i][j-1] {
			return true
		}
	} else if j >= (len(cave_floor[i])-1) {
		if check < cave_floor[i-1][j] &&
			check < cave_floor[i+1][j] &&
			check < cave_floor[i][j-1] {
			return true
		}
	} else {
		if check < cave_floor[i-1][j] &&
			check < cave_floor[i+1][j] &&
			check < cave_floor[i][j+1] &&
			check < cave_floor[i][j-1] {
			return true
		}
	}
	return false
}
