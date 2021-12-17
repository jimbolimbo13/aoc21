package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 20 394 13 sec
// 21 385 27 sec
// 22 395 1 min
// 23 386 2 min
// 24 383 5 min 25
// 25 375 15 minutes
const(
	MAX_DEPTH = 100
	MAX_SCORE = 10000
	// X = 0
	// Y = 1
)

type Integer interface{
	Int() int
}

type chiton interface{
	getCost() int
	getTotal() int
	Int() int
}

type chtn struct {
	cost int
	total int
	dist int
}

func (c chtn) Int() int {
	return c.cost
}

func (c chtn) getCost() int{
	return c.cost
}

func (c chtn) getTotal() int {
	return c.total
}

func Day15() {
	d15_part1()
	d15_part2()
}

func d15_part2(){
	ceiling := parseInput15("input")
	ceiling = quintupleMe(ceiling)

	// for _,v := range ceiling {
	// 	for _, c := range v{
	// 		fmt.Print(" ", c.cost)
	// 	}
	// 	fmt.Println("")
	// }

	total_score := 0
	ceiling = crawlBoard(ceiling)

	total_score = ceiling[0][0].total - ceiling[0][0].cost
	fmt.Println("Day15, Part2:", total_score)
}

func d15_part1() {
	ceiling := parseInput15("input")
	total_score := 0
	ceiling = crawlBoard(ceiling)

	total_score = ceiling[0][0].total - ceiling[0][0].cost

	fmt.Println("Day15, Part1:", total_score)
}

func parseInput15(f string) [][]chtn {
	file, err := os.Open("./data/day15_"+ f +".txt")
	if err != nil {
		fmt.Println(err)
		return make([][]chtn, 0)
	}
	defer file.Close()

	ceiling := make([][]chtn,0)
	//max_xy := []int{0,0}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		new_row := strings.Split(scanner.Text(), "")
		chitons := make([]chtn, len(new_row))
		for i, v := range new_row {
			ch, _ := strconv.Atoi(v)
			chitons[i] = chtn{ch, 0, 0}
		}
		ceiling = append(ceiling, chitons)
	}
	return ceiling
}

func quintupleMe ( c [][]chtn ) [][]chtn {
	larger := make([][]chtn, len(c)*5)


	for i, v := range c {
		for x := 0; x <= 4; x++{
			adder := func (a chtn) chtn {
				a.cost = (a.cost + x)
				if a.cost > 9 { a.cost -= 9 }
				return a
			}
			larger[i + len(c)*x] = mapChtn(adder, v)
		}
	}

	for i, v := range larger {
		larger[i] = append(v, appendChtn(v, 1)...)
	}

	return larger
}

func appendChtn(c []chtn, depth int) []chtn {
	adder := func (a chtn) chtn {
		a.cost = (a.cost + 1)
		if a.cost > 9 { a.cost -= 9 }
		return a
	}
	mapped := mapChtn(adder, c)
	if depth >= 4 {
		return mapped
	} else {
		return append(mapped, appendChtn(mapped, depth +1)...)
	}
}

func mapChtn(f func(chtn) chtn, list []chtn) []chtn{
	mapped := make([]chtn,len(list))
	for i, v := range list{
		mapped[i] = f(v)
	}
	return mapped
}

func scoreCell ( c [][]chtn, crd []int  ) int{
	p_1 := []int{crd[X]+1,crd[Y]}
	p_2 := []int{crd[X], crd[Y]+1}

	dim := []int{0,0}
	dim[X] = len(c[0])
	dim[Y] = len(c)

	// fmt.Println("crd", crd, c[crd[Y]][crd[X]])
	s_1 := MAX_SCORE
	if IsValidIxGrid(p_1[X], p_1[Y], dim) {
		s_1 = 0
		s_1 = c[p_1[Y]][p_1[X]].total
		// fmt.Println("s_1", s_1, c[p_1[Y]][p_1[X]])
	}

	s_2 := MAX_SCORE
	if IsValidIxGrid(p_2[X], p_2[Y], dim) {
		s_2 = 0
		s_2 = c[p_2[Y]][p_2[X]].total
		// fmt.Println("s_2", s_2, c[p_2[Y]][p_2[X]])
	}

	if s_1 == MAX_SCORE && s_2 == MAX_SCORE{
		return c[crd[Y]][crd[X]].getCost()
	} else if s_1 < s_2 {
		// fmt.Println("returning s_1", c[crd[Y]][crd[X]].getCost(), s_1)
		return  c[crd[Y]][crd[X]].getCost() + s_1
	} else {
		// fmt.Println("returning s_2", c[crd[Y]][crd[X]].getCost(), s_2)
		return c[crd[Y]][crd[X]].getCost() + s_2
	}

}

func crawlBoard( c [][]chtn ) [][]chtn {
	start := []int{ len(c[0])-1, len(c)-1 }
	end := []int{0,0}
	next := make([][]int,0)
	new_next := make([][]int,0)
	dim := []int{0,0}
	dim[X] = len(c[0])
	dim[Y] = len(c)

	// 9, 9
	// 9, 8
	// 8, 9
	// 8, 8
	// 9, 7
	// 7, 9
	// 8, 7
	// 7, 8
	//

	next = append(next, start)
	count := 0
	for {
		for _, v := range next{
			c[v[Y]][v[X]].total = scoreCell(c, v)
			// fmt.Println("for loop",c[v[Y]-1][v[X]], v)
			if IsValidIxGrid(v[X], v[Y]-1, dim) && c[v[Y]-1][v[X]].total == 0 {
				new_next = append(new_next, []int{v[X], v[Y]-1})
				c[v[Y]-1][v[X]].total = -1
			}
			if IsValidIxGrid(v[X]-1, v[Y], dim ) && c[v[Y]][v[X]-1].total == 0 {
				new_next = append(new_next, []int{v[X]-1, v[Y]})
				c[v[Y]][v[X]-1].total = -1
			}
		}
		next = new_next
		// fmt.Println(next)
		if c[end[Y]][end[X]].total > 0 { break }
		// if count > 100 {
		// 	fmt.Println(len(next))
		// 	break}
		count ++
	}

	return c

}

// func getManhattan(a []int, b []int) int {
// 	return absInt(a[X] - b[X]) + absInt(a[Y] - b[Y])
// }

// func absInt (x int) int {
// 	if x > 0 {
// 		return x
// 	} else if x < 0 {
// 		return -x
// 	} else {
// 		return 0
// 	}
// }

func scoreForkHalf(c [][]chtn, crd []int, depth int, total int) ([]int, int, int) {
	// TODO refactor for struct
	crd_score := c[crd[Y]][crd[X]].getCost()
	if depth >= MAX_DEPTH { return crd, crd_score, total}
	// fmt.Println(depth, total)
	if depth > 5 && total >= 6*depth  {
		return crd, crd_score, total }
	if total >= 375 {
		return crd, crd_score, total}

	new_total_1 := 0
	new_total_2 := 0
	p_1 := []int{crd[X]+1,crd[Y]}
	p_2 := []int{crd[X], crd[Y]+1}

	sc_1 := MAX_SCORE
	s_1 := MAX_SCORE
	dim := []int{0,0}
	dim[X] = len(c[0])
	dim[Y] = len(c)
	if IsValidIxGrid(p_1[X], p_1[Y], dim) {
		_, s_1, new_total_1 = scoreForkHalf(c, p_1, depth+1, total + c[p_1[Y]][p_1[X]].cost)
		sc_1 = s_1 + c[p_1[Y]][p_1[X]].cost
		// fmt.Println("scores", depth, p_1, c[p_1[Y]][p_1[X]],  s_1 + c[crd[Y]][crd[X]])
		// fmt.Println("scores 1", depth, p_1, c[p_1[Y]][p_1[X]],  s_1 + c[p_1[Y]][p_1[X]])
	}

	sc_2 := MAX_SCORE
	s_2 := MAX_SCORE
	if IsValidIxGrid(p_2[X], p_2[Y], dim) {
		_, s_2, new_total_2= scoreForkHalf(c, p_2, depth+1, total + c[p_2[Y]][p_2[X]].cost)
		sc_2 = s_2 + c[p_2[Y]][p_2[X]].cost
		// fmt.Println("scores", depth, p_2, c[p_2[Y]][p_2[X]], s_2 + c[crd[Y]][crd[X]])
		// fmt.Println("scores 2", depth, p_2, c[p_2[Y]][p_2[X]], s_2 + c[p_2[Y]][p_2[X]])
	}
	// if depth == 0{
	// 	fmt.Println("scores", depth, s_1 + c[p_1[Y]][p_1[X]], s_2 + c[p_2[Y]][p_2[X]])
	// }

	if s_1 == MAX_SCORE && s_2 == MAX_SCORE{
		return crd, c[crd[Y]][crd[X]].cost, total
	} else if sc_1 <= sc_2 {
		return p_1, s_1 + c[crd[Y]][crd[X]].cost, new_total_1
		// return p_1, s_1 + c[p_1[Y]][p_1[X]]
	} else {
		return p_2, s_2 + c[crd[Y]][crd[X]].cost, new_total_2
		// return p_2, s_2 + c[p_2[Y]][p_2[X]]
	}

}

func scoreFork(c [][]chtn, crd []int, depth int) ([]int, int) {
	if depth >= MAX_DEPTH { return crd, c[crd[Y]][crd[X]].cost }

	p_1 := []int{crd[X]+1,crd[Y]}
	p_2 := []int{crd[X], crd[Y]+1}

	sc_1 := MAX_SCORE
	s_1 := MAX_SCORE
	dim := []int{0,0}
	dim[X] = len(c[0])
	dim[Y] = len(c)
	if IsValidIxGrid(p_1[X], p_1[Y], dim) {
		_, s_1 = scoreFork(c, p_1, depth+1)
		sc_1 = s_1 + c[p_1[Y]][p_1[X]].cost
		// fmt.Println("scores", depth, p_1, c[p_1[Y]][p_1[X]],  s_1 + c[crd[Y]][crd[X]])
		// fmt.Println("scores 1", depth, p_1, c[p_1[Y]][p_1[X]],  s_1 + c[p_1[Y]][p_1[X]])
	}

	sc_2 := MAX_SCORE
	s_2 := MAX_SCORE
	if IsValidIxGrid(p_2[X], p_2[Y], dim) {
		_, s_2 = scoreFork(c, p_2, depth+1)
		sc_2 = s_2 + c[p_2[Y]][p_2[X]].cost
		// fmt.Println("scores", depth, p_2, c[p_2[Y]][p_2[X]], s_2 + c[crd[Y]][crd[X]])
		// fmt.Println("scores 2", depth, p_2, c[p_2[Y]][p_2[X]], s_2 + c[p_2[Y]][p_2[X]])
	}
	// if depth == 0{
	// 	fmt.Println("scores", depth, s_1 + c[p_1[Y]][p_1[X]], s_2 + c[p_2[Y]][p_2[X]])
	// }

	if s_1 == MAX_SCORE && s_2 == MAX_SCORE{
		return crd, c[crd[Y]][crd[X]].cost
	} else if sc_1 <= sc_2 {
		return p_1, s_1 + c[crd[Y]][crd[X]].cost
		// return p_1, s_1 + c[p_1[Y]][p_1[X]]
	} else {
		return p_2, s_2 + c[crd[Y]][crd[X]].cost
		// return p_2, s_2 + c[p_2[Y]][p_2[X]]
	}

}

func IsValidIxGrid (x int, y int, dim []int) bool {
	maxX := dim[X]
	maxY := dim[Y]
	if x < maxX && x >=0 && y < maxY && y >= 0 {
		return true
	} else {
		return false
	}
}
