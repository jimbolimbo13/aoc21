package utils

import (
	"fmt"
	"math"
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
	file, err := os.Open("./data/day5_input.txt")
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

	floor = make([][]int, max_xy[1] + 1)
	for i, _ := range floor {
		floor[i] = make([]int, max_xy[0] + 1)
	}
	for _, v := range ortho_lns{
		trace_vec(v)
	}

	// for _, v := range floor {
	// 	fmt.Println(v)
	// }
	fmt.Println("Day5, Part1:",count_overlap(2))
}

func count_overlap(threshold int) int {
	count := 0
	for _, y := range floor {
		for _, xy := range y {
			if xy >= threshold{
				count += 1
			}
		}
	}
	return count
}

func trace_vec(vec []int){
	x_diff := int(math.Abs(float64(vec[0] - vec[2])))
	y_diff := int(math.Abs(float64(vec[1] - vec[3])))
	run := greater(x_diff, y_diff)
	for i := vec[(lesser(vec[run%2], vec[(run%2)+2])*2)+run];
		i <= vec[(greater(vec[run%2], vec[(run%2)+2])*2)+run];
		i ++ {
			x:=0
			y:=0

			if run == 0{
				x = i
				y = vec[1]
			} else {
				x = vec[0]
				y = i
			}

			floor[y][x] += 1
	}
}

func lesser(a int, b int) int{
	if a < b {
		return 0
	} else {
		return 1
	}
}

func greater(a int, b int) int{
	if a >= b {
		return 0
	} else {
		return 1
	}
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
