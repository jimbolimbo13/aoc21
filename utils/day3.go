package utils

import (
	"bufio"
	// "bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day3() {
	part1()
	part2()
}

func part2() {

	file, err := os.Open("./data/day3_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	first := scanner.Text()
	first_split := strings.Split(first, "")
	width := len(first)
	total := 0
	ones := make([][]string, 0)
	zeros := make([][]string, 0)

	if first_split[0] == "1" {
		ones = append(ones, first_split)
	} else {
		zeros = append(zeros, first_split)
	}
	total += 1

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "")
		if split[0] == "1" {
			ones = append(ones, split)
		} else {
			zeros = append(zeros, split)
		}
		total += 1
	}

	o2_loop := ones
	co2_loop := zeros
	if len(zeros) > len(ones) {
		o2_loop = zeros
		co2_loop = ones
	}

	for i := 1; i < width; i++ {
		o2_ones := make([][]string, 0)
		o2_zeros := make([][]string, 0)
		co2_ones := make([][]string, 0)
		co2_zeros := make([][]string, 0)

		main_loop := 0
		// len_dif := 0
		longer := ""
		if len(co2_loop) > len(o2_loop) {
			main_loop = len(o2_loop)
			longer = "co2"
			// len_dif = (co2_loop) - len(o2_loop)
		} else {
			main_loop = len(co2_loop)
			longer = "o2"
		}

		for j := 0; j < main_loop; j++ {
			if o2_loop[j][i] == "1" {
				o2_ones = append(o2_ones, o2_loop[j])
			} else {
				o2_zeros = append(o2_zeros, o2_loop[j])
			}
			if co2_loop[j][i] == "1" {
				co2_ones = append(co2_ones, co2_loop[j])
			} else {
				co2_zeros = append(co2_zeros, co2_loop[j])
			}
		}

		remaining := o2_loop[main_loop:]
		if longer == "co2" {
			remaining = co2_loop[main_loop:]
		}

		for j := 0; j < len(remaining); j ++{
			if longer == "o2" {
				if remaining[j][i] == "1" {
				o2_ones = append(o2_ones, remaining[j])
				} else {
				o2_zeros = append(o2_zeros, remaining[j])
				}
			} else {
				if remaining[j][i] == "1" {
					co2_ones = append(co2_ones, remaining[j])
				} else {
					co2_zeros = append(co2_zeros, remaining[j])
				}
			}
		}

		if len(o2_zeros) > len(o2_ones) {
			o2_loop = o2_zeros
		} else {
			o2_loop = o2_ones
		}

		if len(co2_ones) == len(co2_zeros){
			co2_loop = co2_zeros
		} else if len(co2_ones) > len(co2_zeros) {
			if len(co2_zeros) == 0 {
				co2_loop = co2_ones
			} else {
				co2_loop = co2_zeros
			}
		} else {
			if len(co2_ones) == 0 {
				co2_loop = co2_zeros
			} else {
				co2_loop = co2_ones
			}
		}
	}


	o2_factor, _ := strconv.ParseInt(strings.Join(o2_loop[0],""), 2, 64)
	co2_factor, _ := strconv.ParseInt(strings.Join(co2_loop[0],""), 2, 64)
	fmt.Println("Day3, Part2:", o2_factor * co2_factor)
}


func part1() {
	file, err := os.Open("./data/day3_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	first := scanner.Text()
	first_split := strings.Split(first, "")
	width := len(first)
	bits := make([]int, width)
	total := 0

	for i := 0; i < width; i ++ {
		next_bit, _ := strconv.Atoi(first_split[i])
		bits[i] += next_bit
	}
	total += 1

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "")
		for i := 0; i < width; i ++ {
			next_bit, _ := strconv.Atoi(split[i])
			bits[i] += next_bit
		}
		total += 1
	}

	gamma_val := 0
	eps_val := 0
	for i := 0; i < width; i ++ {
		if bits[i] > total / 2 {
			gamma_val += 1 << (width-1-i)
		} else {
			eps_val += 1 << (width-1-i)
		}
	}

	fmt.Println("Day3, Part1:", gamma_val * eps_val)
}
