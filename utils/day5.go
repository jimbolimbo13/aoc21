package utils

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

var(
	floor = make([][]int, 0)
)
func Day5() {
	d5_part1()
	// d5_part2()
}

func d5_part2(){
	fmt.Println("Day5, Part2")
}

func d5_part1() {
	file, err := os.Open("./data/day5_example.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ortho_lns := make([][]int, 0)
	line := make([]string, 4)
	line_int := make([]int, 4)
	max_xy := []int{0,0}
	for scanner.Scan() {
		split := strings.Fields(scanner.Text())
		first := strings.Split(split[0], ",")
		second := strings.Split(split[2], ",")
		if first[0] == second [0] || first[1] == second[1]{
			line = []string{first[0], first[1], second[0], second[1]}
			line_int = str_to_int(line)
			ortho_lns = append(ortho_lns, line_int)
			mag_max, dir_max := get_max_vec(line_int)
			if mag_max > max_xy[dir_max] {
				max_xy[dir_max] = mag_max
			}
		}
	}


	fmt.Println(max_xy)
	fmt.Println(ortho_lns)
	fmt.Println("Day5, Part1:")
}

func get_max_vec(vec []int) (int, int){
	ret_mag := 0
	ret_dir := 0  // 0 for x, 1 for y
	for i, v := range vec {
		if v >= ret_mag {
			ret_mag = v
			ret_dir = i%2
		}
	}
	return ret_mag, ret_dir
}

func str_to_int(str []string) []int {
	ret_int := make([]int, len(str))
	for i, v := range str{
		ret_int[i], _ = strconv.Atoi(v)
	}
	return ret_int
}
