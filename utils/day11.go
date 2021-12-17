package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"math"
	"time"
)

var(
	run_total = 100
	total_flashes = 0
	run_number = 0
	all_flashed = false
)

func Day11() {
	flash_graphics()
	// d11_part1()
	// d11_part2()
}

func flash_graphics(){
	file, err := os.Open("./data/day11_input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	octopi := make([]int,0,100)
	dim := make([]int,2)
	ix := 0

	for scanner.Scan() {
		new_row := strings.Split(scanner.Text(), "")
		dim[1] = len(new_row)
		for _, v := range new_row {
			nrg, _ := strconv.Atoi(v)
			octopi = append(octopi, nrg)
			ix ++
		}
		dim[0] ++
	}

	for !all_flashed {
		run_step(&octopi, dim)
		print_octo_flash(octopi, dim)
		// time.Sleep(550 * time.Millisecond)
	}
	run_step(&octopi, dim)
	print_octo_flash(octopi, dim)

	fmt.Println()

}

func print_octo_flash(octopi []int, dim []int){
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
        cmd.Run()
		Print_grid_flash(octopi, dim)
}

func d11_part2(){
	file, err := os.Open("./data/day11_input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	octopi := make([]int,0,100)
	dim := make([]int,2)
	ix := 0

	for scanner.Scan() {
		new_row := strings.Split(scanner.Text(), "")
		dim[1] = len(new_row)
		for _, v := range new_row {
			nrg, _ := strconv.Atoi(v)
			octopi = append(octopi, nrg)
			ix ++
		}
		dim[0] ++
	}

	run_count := 0
	for !all_flashed {
		run_step(&octopi, dim)
		run_count ++
	}

	fmt.Println("Day11, Part2:", run_count)
}

func d11_part1() {
	file, err := os.Open("./data/day11_input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	octopi := make([]int,0,100)
	dim := make([]int,2)
	ix := 0

	for scanner.Scan() {
		new_row := strings.Split(scanner.Text(), "")
		dim[1] = len(new_row)
		for _, v := range new_row {
			nrg, _ := strconv.Atoi(v)
			octopi = append(octopi, nrg)
			ix ++
		}
		dim[0] ++
	}

	for i:=0; i< run_total; i++{
		run_step(&octopi, dim)
	}

	fmt.Println("Day11, Part1:", total_flashes)
}

func run_step(octopi *[]int, dim []int){
	step1(octopi)
	for {
		over_9 := func(i int) bool { return i > 9}
		flashers := SelectIntWithIx(over_9, *octopi)
		if len(flashers) == 0 { break }
		for i, _ := range flashers{
			octo_flash(i, octopi, dim)
		}
		time.Sleep(30 * time.Millisecond)
		print_octo_flash(*octopi, dim)
	}
	// reset flashed
	negative := func(i int) bool { return i < 0 }
	flashed := SelectIntWithIx(negative, *octopi)
	for i := range flashed {
		(*octopi)[i] = 0
	}
	if len(flashed) == (dim[0]*dim[1]){
		all_flashed = true
	}

}

func octo_flash(o int, list *[]int, dim []int){
	// 1 2 3
	// 4 5 6
	// 7 8 9
	//
	//5-row+1  5-row  5-row-1
	//5-1       5+1
	//5+row-1  5+row  5+row+1
	ix := 0
	for i := -1; i <= 1; i++{
		for j := -1; j <= 1; j++{
			if j == 0 && i == 0 {continue}
			ix = o + (i*dim[1]) + (1*j)
			if Is_valid_ix(o, ix, dim){
				// fmt.Println(ix,(*list)[ix])
				(*list)[ix] = (*list)[ix] + 1
			}
		}

	}
	(*list)[o] = -100
	total_flashes ++
}

func Is_valid_ix(start int, next int, dim []int) bool {
	s_row := start / dim[0]
	s_col := start % dim[1]
	n_row := next / dim[0]
	n_col := next % dim[1]

	// fmt.Println("valid",start,next,s_row,s_col,n_row,n_col)

	if next < 0 || next >= dim[0] * dim[1] {
		return false
	} else if math.Abs(float64(s_row - n_row)) > 1 || math.Abs(float64(s_col - n_col)) > 1{
		return false
	} else if math.Abs(float64(s_row - n_row)) == 0 && math.Abs(float64(s_col - n_col)) == 0{
		return false
	} else  {
		return true
	}
}

func step1(octopi *[]int) {
	increment := func(i int) int {return i+1}
	*octopi = MapInt(increment, *octopi)
}

func MapInt(f func(int) int, list []int) []int{
	mapped := make([]int,len(list))
	for i, v := range list{
		mapped[i] = f(v)
	}
	return mapped
}

func SelectIntWithIx(f func(int) bool, list []int) map[int]int {
	selected := make(map[int]int)
	for i, v := range list{
		if f(v){
			selected[i] = v
		}
	}
	return selected
}

func Print_grid(octopi []int, dim []int) {
	for i := 0; i < dim[0]; i++{
		for j := 0; j < dim[1]; j++{
			spacer := "  "
			if octopi[j +(i*dim[1])] > 99 {
				spacer = ""
			} else if octopi[j +(i*dim[1])] > 9 {
				spacer = " "
			}
			fmt.Print(octopi[j +(i*dim[1])], spacer)
			// fmt.Print(j +(i*dim[1]))
		}
		fmt.Println()
	}
}

func Print_grid_flash(octopi []int, dim []int) {
	for i := 0; i < dim[0]; i++{
		for j := 0; j < dim[1]; j++{
			spacer := " "
			sign := ""
			switch octopi[j +(i*dim[1])]{
				case 1:
					sign = "･"
				case 2:
					sign = "･"
				case 3:
					sign = "･"
				case 4:
					sign = "𐄁"
				case 5:
					sign = ""
				case 6:
					sign = ""
				case 7:
					sign = ""
				case 8:
					sign = ""
				case 9:
					sign = ""
				// case 0:
				// 	sign = ""
				// 	spacer = ""
			}
			switch {
				case octopi[j +(i*dim[1])] > 9:
					sign = ""
					spacer = ""
				case octopi[j +(i*dim[1])] <= 0:
					sign = " "
			}
			fmt.Printf("%s%s%s", spacer, sign, spacer)
			// fmt.Print(j +(i*dim[1]))
		}
		fmt.Println()
	}
}
