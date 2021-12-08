package utils

import (
	"fmt"
	"math"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sync"
)

var(
	floor = make([][]int, 0)
	floor_wait sync.WaitGroup
	floor_lock sync.Mutex
)
func Day5() {
	d5_part1()
	d5_part2()
}

func d5_part2(){
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

		line = []string{first[0], first[1], second[0], second[1]}
		line_int = str_to_int(line)
		ortho_lns = append(ortho_lns, line_int)
		mag_max, dir_max := get_max_vec(line_int)
		if dir_max < 0 {
			if mag_max > max_xy[0]{
				max_xy[0] = mag_max
			}
			if mag_max > max_xy[1]{
				max_xy[1] = mag_max
			}
		} else {
			if mag_max > max_xy[dir_max] {
				max_xy[dir_max] = mag_max
			}
		}
	}

	floor = make([][]int, max_xy[1] + 2)
	for i := range floor {
		floor[i] = make([]int, max_xy[0] + 2)
	}
	for _, v := range ortho_lns{
		floor_wait.Add(1)
		go trace_vec(v)
	}
	floor_wait.Wait()

	// for _, v := range floor {
	// 	fmt.Println(v)
	// }
	fmt.Println("Day5, Part2:",count_overlap(2))
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
			if dir_max < 0 {
				if mag_max > max_xy[0]{
					max_xy[0] = mag_max
				}
				if mag_max > max_xy[1]{
					max_xy[1] = mag_max
				}
			} else {
				if mag_max > max_xy[dir_max] {
					max_xy[dir_max] = mag_max
				}
			}
		}
	}

	floor = make([][]int, (max_xy[0] + 1))
	for i := range floor {
		floor[i] = make([]int, (max_xy[1] + 1))
	}
	// floor_wait.Add(len(ortho_lns))
	for _, v := range ortho_lns{
		floor_wait.Add(1)
		go trace_vec(v)
	}
	floor_wait.Wait()

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
	defer floor_wait.Done()
	x_diff := vec[2] - vec[0]
	y_diff := vec[3] - vec[1]
	dir_xy := []int{1,1}
	if x_diff < 0{
		dir_xy[0] = -1
	}
	if y_diff < 0{
		dir_xy[1] = -1
	}

	diff_xy := []int{int(math.Abs(float64(x_diff))),
		int(math.Abs(float64(y_diff)))}
	run := greater(diff_xy[0], diff_xy[1])
	diag := false
	if diff_xy[0] == diff_xy[1] {
		diag = true
	}
	xy := []int{vec[0],vec[1]}
	for i := 0; i <= diff_xy[run]; i ++ {
		// if xy[0] >= 989 || xy [1] >= 989 {
		// 	fmt.Println(vec,run,diag,xy) //,len(floor[xy[0]]),len(floor[xy[1]]))
		// }
		floor_lock.Lock()
		floor[xy[0]][xy[1]] += 1
		floor_lock.Unlock()
		if diag {
			xy[0] += (1*dir_xy[0])
			xy[1] += (1*dir_xy[1])
		} else {
			xy[run] += (1*dir_xy[run])
		}

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
	if ret_mag == vec[1^ret_dir] || ret_mag == vec[1^ret_dir + 2]{
		ret_dir = -1
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
