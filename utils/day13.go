package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	X = 0
	Y = 1
)

func Day13() {
	// d13_part1()
	d13_part2()
}

func d13_part2(){
	file, err := os.Open("./data/day13_input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	max_xy := []int{0,0}
	lines := make([][]int,0)
	read_coords := true
	commands := make([][]int,0)

	for scanner.Scan() {
		new := strings.Split(scanner.Text(), ",")
		if read_coords && len(new) == 2{
			new_int := SliceStrToSliceInt(new)
			mag_max, dir_max :=	Get_max_vec(new_int)
			if mag_max > max_xy[dir_max] {
				max_xy[dir_max] = mag_max
			}
			lines = append(lines, new_int)
		} else if len(new[0]) > 1{
			new = strings.Fields(scanner.Text())
			command := strings.Split(new[2], "=")
			if command[0] == "x"{
				coord,_ := strconv.Atoi(command[1])
				commands = append(commands, []int{X,coord})
			} else {
				coord,_ := strconv.Atoi(command[1])
				commands = append(commands, []int{Y,coord})
			}
		} else if len(new) == 1 {
			read_coords = false
		}
	}


	dots := make([][]int,max_xy[Y]+1)
	for i := range dots {
		dots[i] = make([]int,max_xy[X]+1)
	}
	for _,v := range lines{
		dots[v[Y]][v[X]] = 1
	}
	flip_zero := func(a int) int {
		if a == 0 {
			return 1
		} else {
			return 9
		}
	}

	for i := range commands{
		fmt.Print("fold ",i)
		if i == 100 {

			for _,v := range dots{
				// fmt.Println(v)
				Print_grid_flash(MapInt(flip_zero, v), []int{1,len(v)})
			}
		}
		// if i >= 8 {
		// 	if commands[i][0] == X {
		// 		PrintColumnInt(dots, commands[i][1])
		// 	} else {
		// 		fmt.Println(dots[commands[i][1]])
		// 	}
		// }
		// dots , max_xy = foldGrid(dots, commands[i][1], commands[i][0], []int{len(dots),len(dots[0])})
		dots , max_xy = foldGrid(dots, commands[i][1], commands[i][0], max_xy)
		fmt.Println(len(dots), len(dots[0]),max_xy)
	}

	for _,v := range dots{
		// fmt.Println(v)
		Print_grid_flash(MapInt(flip_zero, v), []int{1,len(v)})
	}
	fmt.Println("Day13, Part2:")
}

func d13_part1() {
	file, err := os.Open("./data/day13_input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	max_xy := []int{0,0}
	lines := make([][]int,0)
	read_coords := true
	commands := make([][]int,0)

	for scanner.Scan() {
		new := strings.Split(scanner.Text(), ",")
		if read_coords && len(new) == 2{
			new_int := SliceStrToSliceInt(new)
			mag_max, dir_max :=	Get_max_vec(new_int)
			if mag_max > max_xy[dir_max] {
				max_xy[dir_max] = mag_max
			}
			lines = append(lines, new_int)
		} else if len(new[0]) > 1{
			new = strings.Fields(scanner.Text())
			command := strings.Split(new[2], "=")
			if command[0] == "x"{
				coord,_ := strconv.Atoi(command[1])
				commands = append(commands, []int{X,coord})
			} else {
				coord,_ := strconv.Atoi(command[1])
				commands = append(commands, []int{Y,coord})
			}
		} else if len(new) == 1 {
			read_coords = false
		}
	}


	dots := make([][]int,max_xy[Y]+1)
	for i := range dots {
		dots[i] = make([]int,max_xy[X]+1)
	}
	for _,v := range lines{
		dots[v[Y]][v[X]] = 1
	}

	// for _,v := range dots{
	// 	fmt.Println(v)
	// }

	dots , max_xy = foldGrid(dots, commands[0][1], commands[0][0], max_xy)

	total_dots := 0
	for _,v := range dots{
		for _,k := range v{
			total_dots += k
		}
	}


	fmt.Println("Day13, Part1:",total_dots)
}

func foldGrid(grid [][]int, coord int, dir int, dim []int) ([][]int, []int) {
	// new_x := 0
	// new_y := 1
	fmt.Println(" ",coord,dir,dim, len(grid),len(grid[0]))
	folded := make([][]int,0)
	comp  := func(a int, b int) int{
		if a == 1 || b == 1{
			return 1
		} else {
			return 0
		}
	}

	if dir == Y {
		folded = make([][]int,coord)
		if coord < 40 {
			fmt.Println("lengths",len(grid[0]), len(grid[dim[Y]-0]))
		}
		for i:=0; i<coord; i++{
			folded[i] = mergeSliceInt(grid[i], grid[dim[Y]-i], comp)
		}

	} else {
		folded = make([][]int,len(grid))
		if coord < 100{
			fmt.Println("lengths",len(grid[0][:coord]),len(grid[0][coord+1:]))
		}
		for i,v := range grid{
			folded[i] = mergeSliceInt(v[:coord], ReverseInt( v[coord+1:]), comp)
		}
	}
	dim[Y] = len(folded) -1
	dim[X] = len(folded[0]) -1
	return folded, dim
}

func mergeSliceInt (s1 []int, s2 []int, f func(a int, b int) int) []int {
	merged := s1
	for i,v := range merged {
		merged[i] = f(v, s2[i])
	}
	return merged
}

func ReverseInt(s []int) []int{
	reversed := s
	last := 0
	for i:=0; i<=(len(s)/2); i++{
		last = reversed[len(s)-i-1]
		reversed[len(s)-i-1] = reversed[i]
		reversed[i] = last
	}
	return reversed
}

func PrintColumnInt(g [][]int, col int){

	for _,v := range g {
		fmt.Println(v[col])
	}
}
