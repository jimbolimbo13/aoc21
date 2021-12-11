package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day11() {
	d11_part1()
	d11_part2()
}

func d11_part2(){
	fmt.Println("Day11, Part2")
}

func d11_part1() {
	file, err := os.Open("./data/day11-example.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	octopi := make([]int,0,100)
	octo_nrg := new(map[int][]int)
	dim := make([]int,2)

	for scanner.Scan() {
		new_row := strings.Split(scanner.Text(), "")
		dim[1] = len(new_row)
		for _, v := range new_row {
			nrg, _ := strconv.Atoi(v)
			octopi = append(octopi, nrg)
			if _, ok := octo_nrg[nrg]; ok {
				octo_nrg = append(octo_nrg[nrg])
			}
			// else {
			// 	octo_nrg[nrg] = []int{(len(octopi)-1)}
			// }
		}
	}
	fmt.Println(octopi)
	fmt.Println("Day11, Part1:")
}
