package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day13() {
	d13_part1()
	d13_part2()
}

func d13_part2(){
	fmt.Println("Day13, Part2:")
}

func d13_part1() {
	file, err := os.Open("./data/day13_example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	max_xy := []int{0,0}
	lines := make([][]int,0)

	for scanner.Scan() {
		new := strings.Split(scanner.Text(), ",")
		new_int := SliceStrToSliceInt(new)
		mag_max, dir_max :=	Get_max_vec(new_int)
		if mag_max > max_xy[dir_max] {
			max_xy[dir_max] = mag_max
		}
		lines = append(lines, new_int)
	}

	dots := make([][]int,max_xy[1]+1)
	for i := range dots {
		dots[i] = make([]int,max_xy[0]+1)
	}
	for _,v := range lines{
		dots[v[1]][v[0]] = 1
	}

	for _,v := range dots{
		fmt.Println(v)
	}
	fmt.Println("Day13, Part1:")
}
