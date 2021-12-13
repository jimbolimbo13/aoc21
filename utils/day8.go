package utils

import (
	"bufio"
	"fmt"
	"os"
<<<<<<< HEAD
)

func DayXX() {
	dXX_part1()
	dXX_part2()
}

func dXX_part2(){
	fmt.Println("DayXX, Part2")
}

func dXX_part1() {
	file, err := os.Open("./data/dayXX_input.txt")
=======
	"strings"
	"math"
	// "regexp"
)

func Day8() {
	d8_part1()
	d8_part2()
}

func d8_part2(){
	file, err := os.Open("./data/day8_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	before_split := make([][]string, 0)
	after_split := make([][]string, 0)
	// digit_counts := make([]int, 10)

	for scanner.Scan() {
		new_line := strings.Split(scanner.Text(), "|")
		before_split = append(before_split, strings.Fields(new_line[0]))
		after_split = append(after_split, strings.Fields(new_line[1]))
	}

	// top, up-r, up-l, mid, low-r, low-l, bottom
	// with 1 and 7, we can determine top and (u-r / l-r)
	// top
	// (u-r / l-r)
	// with 4 and 8, we can get (bot / l-l)
	// (bot / l-l)
	// u-l and mid are the parts of 4 that aren't (u-r / l-r)
	// (u-l / mid)
	// check the length 6 items
	// 0 is missing (u-l / mid) - diff with 8 gives us mid
		// mid
		// u-l
	// 9 is missing (l-l / bot) - diff with 8 gives l-l
		// l-l
		// bot
	// 6 is missing (u-r / l-r) - diff with 8 gives u-r
		// u-r
		// l-r
	// no need to check length 5 items
	// 3 is missing l-l and (u-l / mid) - diff with 8 and l-l gives u-l
	// 2 is missing
	//
	// functions:
	// parse lines (takes both splits)
	// string diff (takses two strings, returns what's not common)
	final_count := 0
	for i, _ := range before_split {
		line_count := parse_line(before_split[i], after_split[i])
		// fmt.Println(line_count)
		final_count += line_count
	}
	// parse_line(before_split[0], after_split[0])
	// for _, output := range after_split {
	// 	for _, v := range output{
	// 		switch len(v) {
	// 		case 2, 3, 4, 7:
	// 			digit_counts += 1
	// 		}
	// 	}
	// }
	fmt.Println("Day8, Part2", final_count)
}

func parse_line (before []string, after []string) int {

	ten_digits := make([]string, 4, 10) // 1, 7, 4, 8, everything
	for _, v := range before{
		switch len(v) {
		case 2: // 1
			ten_digits[0] = v
		case 3: // 7
			ten_digits[1] = v
		case 4: // 4
			ten_digits[2] = v
		case 7: // 8
			ten_digits[3] = v
		default:
			ten_digits = append(ten_digits, v)
		}
	}


	segments := make([]string, 7) // top, up-r, up-l, mid, low-r, low-l, bottom
								//   0	  1 	2 	  3 	4 	  5 	6
								//   0: 21-3
								//   2: 21-2-4
								//   3: 21-2-5
								//   5: 21-1-5
								//   6: 21-1
								//   9: 21-5
	ur_lr := ""
	bot_ll := ""
	ul_mid := ""

	segments[0] = string_diff(ten_digits[0], ten_digits[1])
	ur_lr = string_diff(segments[0], ten_digits[0])
	d4_8 := string_diff(ten_digits[2], ten_digits[3])
	bot_ll = string_diff(d4_8, segments[0])
	ul_mid = string_diff(ten_digits[2], ur_lr)

	for i := 4; i < 10; i++ {
		if len(ten_digits[i]) == 6{
			missing := string_diff(ten_digits[3], ten_digits[i])
			if str_contains(missing, ul_mid) {
				segments[3] = missing
				segments[2] = string_diff(missing, ul_mid)
			} else if str_contains(missing, bot_ll) {
				segments[5] = missing
				segments[6] = string_diff(missing, bot_ll)
			} else if str_contains(missing, ur_lr){
				segments[1] = missing
				segments[4] = string_diff(missing, ur_lr)
			}
		}
	}

	// fmt.Println(segments)

	seg_map := make(map[string]int)
	for i, v := range segments {
		seg_map[v] = i
	}

	counts := 0
	multi := 0
	for i, v := range after {
		multi = 1000 / int(math.Pow10(i))
		// fmt.Println(multi)
		counts += (read_output(seg_map, v) * multi)
	}

	return counts
}

func read_output(seg_map map[string]int, find string) int {
	switch len(find) {
	case 2: // 1
		return 1
	case 3: // 7
		return 7
	case 4: // 4
		return 4
	case 7: // 8
		return 8
	default:
		return parse_segments(seg_map, find)
	}
}

func parse_segments(seg_map map[string]int, find string) int {
	seg_value := 0
	has_two := false
	for _, v := range find {
		seg_value += seg_map[string(v)]
		if seg_map[string(v)] == 2{
			has_two = true
		}
	}

	value := 0
	switch seg_value {
	case 21-3:
		return 0
	// case 21-2-4:
	// 	return 2
	case 21-2-5:
		return 3
	case 21-1-5:
		if has_two {
			return 5
		} else{
			return 2
		}
	case 21-1:
		return 6
	case 21-5:
		return 9
	}
	//   0: 21-3
	//   2: 21-2-4
	//   3: 21-2-5
	//   5: 21-1-5
	//   6: 21-1
	//   9: 21-5
	return value
}

func str_contains(find string, within string) bool {
	for _, v := range within {
		if find == string(v) {
			return true
		}
	}
	return false
}

func string_diff(sta string, stb string) string {
	longer := stb
	shorter := sta
	diff := ""
	if len(sta) > len(stb) {
		longer = sta
		shorter = stb
	}
	map_shorter := make(map[rune]bool)
	for _, v := range shorter {
		map_shorter[v] = true
	}

	// use a map for the comparison, loop
	for _, v := range longer {
		if _, ok := map_shorter[v]; ok{
			continue
		} else {
			diff += string(v)
		}
	}
	return diff
}

func d8_part1() {
	file, err := os.Open("./data/day8_input.txt")
>>>>>>> ca33d39433c2a49dee324157d8f7be5305c8481e
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

<<<<<<< HEAD
	for scanner.Scan() {
		// process data
	}
	fmt.Println("DayXX, Part1:")
=======
	before_split := make([][]string, 0)
	after_split := make([][]string, 0)
	// digit_counts := make([]int, 10)
	digit_counts := 0

	for scanner.Scan() {
		new_line := strings.Split(scanner.Text(), "|")
		before_split = append(before_split, strings.Fields(new_line[0]))
		after_split = append(after_split, strings.Fields(new_line[1]))
	}

	for _, output := range after_split {
		for _, v := range output{
			switch len(v) {
			case 2, 3, 4, 7:
				digit_counts += 1
			}
		}
	}
	fmt.Println("Day8, Part1:", digit_counts)
>>>>>>> ca33d39433c2a49dee324157d8f7be5305c8481e
}
