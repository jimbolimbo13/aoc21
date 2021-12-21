package utils

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func Day18() {
	// d18_part1()
	d18_part2()
}

func d18_part2(){
	snailNums := parseInput18("input")

	maxMag := 0
	for i := range snailNums {
		for j := range snailNums{
			if i == j { continue }
			next_i := snailNums[i].deepCopy()
			// fmt.Println(next_i.toString())
			next_j := snailNums[j].deepCopy()
			// fmt.Println(next_j.toString())
			next := addNum(next_i, next_j)
			// fmt.Println(next.toString())
			next.reduce()
			newMag := next.getMag()
			// fmt.Println("new Mag", newMag)
			if newMag > maxMag {
				maxMag = newMag
				// fmt.Println("new Max:", maxMag, snailNums[i].toString(),"+", snailNums[j].toString())
			}
		}

	}
	fmt.Println("Day18, Part2:", maxMag)
}

func d18_part1() {
	snailNums := parseInput18("input")

	var total *number
	for i, v := range snailNums {
		if i == 0 {
			total = v
			continue
		}
		// fmt.Println("adding", total.toString(), "+", v.toString())
		total = addNum(total, v)
		// fmt.Println("before reduce", total.toString())
		total.reduce()
		// fmt.Println("after reduct", total.toString())
		// if i == 4 { break }
	}

	// fmt.Println(total.toString())

	fmt.Println("Day18, Part1:", total.getMag())
}

func seekDepth( root *number , depth int) (*number, int) {
	// if depth  >= 4 { return root, depth}
	var leftNum *number
	leftDepth := depth + 1
	if root.leftNum != nil {
		leftNum, leftDepth = seekDepth(root.leftNum, leftDepth)
	}
	var rightNum *number
	rightDepth := depth + 1
	if root.rightNum != nil {
		rightNum, rightDepth = seekDepth(root.rightNum, rightDepth)
	}
	if leftNum == nil && rightNum == nil{
		return root, depth
	} else if leftNum != nil && rightNum == nil {
		return  leftNum, leftDepth
	} else  if leftNum == nil && rightNum != nil{
		return rightNum, rightDepth
	}
	if leftDepth >= rightDepth {
		return leftNum, leftDepth
	} else {
		return rightNum, rightDepth
	}
}

func seekSplit( root *number ) (*number) {
	// return condition
	if root.leftNum == nil && root.rightNum == nil{
		if root.leftVal.val > 9 || root.rightVal.val > 9 {
			return root
		} else {
			return nil
		}
	}

	consider := (root.leftVal.val > 9 || root.rightVal.val > 9)
	var leftNum *number
	if root.leftNum != nil {
		leftNum = seekSplit(root.leftNum)
	}
	var rightNum *number
	if root.rightNum != nil {
		rightNum = seekSplit(root.rightNum)
	}

	if leftNum == nil && consider{
		return root
	} else if leftNum != nil && rightNum == nil {
		return  leftNum
	} else  if leftNum == nil && rightNum != nil{
		return rightNum
	} else {
		return leftNum
	}
}

//CHANGEME
func newNum() *number{
	retNum := number{
		leftVal:  numVal{},
		rightVal: numVal{},
		parent:   nil,
		leftNum:  nil,
		rightNum: nil,
	}
	return &retNum
}

func addNum (left *number, right *number) *number {
	sum := number{
		leftVal:  numVal{},
		rightVal: numVal{},
		parent:   nil,
		leftNum:  left,
		rightNum: right,
	}
	sum.leftNum.parent = &sum
	sum.rightNum.parent = &sum
	return &sum
}

type number struct{
	leftVal numVal
	rightVal numVal
	parent *number
	leftNum *number
	rightNum *number
	// nestLvl int
}

type numVal struct {
	val int
}

func (n *numVal) add(v int) {
	n.val += v
}

func (n *number) getRight() *numVal {
	if n.parent == nil {
		return nil
	} else
	if n.parent.rightNum != nil && n.parent.rightNum == n {
		return n.parent.getRight()
	} else if n.parent.rightNum != nil {
		if n.parent.rightNum.leftNum == nil {
			return &n.parent.rightNum.leftVal
		} else {
			return n.parent.rightNum.leftNum.diveLeft()
		}
	} else if n.parent.rightNum == nil {
		return &n.parent.rightVal
	}
	return nil
}

func (n *number) getLeft() *numVal {
	if n.parent == nil {
		return nil
	} else if n.parent.leftNum != nil && n.parent.leftNum == n {
		return n.parent.getLeft()
	} else if n.parent.leftNum != nil {
		if n.parent.leftNum.rightNum == nil {
			return &n.parent.leftNum.rightVal
		} else {
			return n.parent.leftNum.rightNum.diveRight()
		}
	} else if n.parent.leftNum == nil {
		return &n.parent.leftVal
	}
	return nil
}

func (n *number) diveLeft() *numVal {
	if n.leftNum == nil {
		return &n.leftVal
	} else {
		return n.leftNum.diveLeft()
	}
}

func (n *number) diveRight() *numVal {
	if n.rightNum == nil {
		return &n.rightVal
	} else {
		return n.rightNum.diveRight()
	}
}

func (n number) print() {
	fmt.Print("[")
	if n.leftNum == nil {
		fmt.Print(n.leftVal.val, ",")
	} else {
		n.leftNum.print()
		fmt.Print(",")
	}
	if n.rightNum == nil {
		fmt.Print(n.rightVal.val)
	} else {
		n.rightNum.print()
	}
	fmt.Print("]")
}

func (n number) toString() string {
	retS := "["
	if n.leftNum == nil {
		retS += (strconv.Itoa(n.leftVal.val) + ",")
	} else {
		retS += n.leftNum.toString() + ","
	}
	if n.rightNum == nil {
		retS += strconv.Itoa(n.rightVal.val)
	} else {
		retS += n.rightNum.toString()
	}
	retS += "]"
	return retS
}

func (n number) getReduced() number {
	n.reduce()
	return n
}

func (n number) reduce() {
	checkDeep := true
	checkSplit := true

	for {
		if !checkDeep && !checkSplit { break }
		for checkDeep {
			if checkDeep {
				deepest, depth := seekDepth(&n, 0)
				if depth >= 4 {
					// fmt.Println("exploding", deepest.toString())
					deepest.explode()
					// fmt.Println("exploded:", n.toString())
					checkSplit = true
				} else {
					checkDeep = false
				}
			}
		}
		if checkSplit {
			splitn := seekSplit(&n)
			if splitn != nil{
				// fmt.Println("splitting", splitn.toString())
				splitn.split()
				// fmt.Println("splitted:", n.toString())
				checkDeep = true
			} else {
				checkSplit = false
			}
		}
	}
}

func (n *number) split() {
	if n.leftVal.val > 9 {
		newLeft := n.leftVal.val / 2
		newRight := newLeft
		rem := math.Remainder( float64(n.leftVal.val), 2.0)
		if rem != 0 {
			newRight ++
		}
		temp := newNum()
		temp.leftVal.val = newLeft
		temp.rightVal.val = newRight
		n.leftVal.val = 0
		n.leftNum = temp
		temp.parent = n
	} else if n.rightVal.val > 9 {
		newLeft := n.rightVal.val / 2
		newRight := newLeft
		rem := math.Remainder( float64(n.rightVal.val), 2.0)
		if rem != 0 {
			newRight ++
		}
		temp := newNum()
		temp.leftVal.val = newLeft
		temp.rightVal.val = newRight
		n.rightVal.val = 0
		n.rightNum = temp
		temp.parent = n
	}
}

func (n *number) explode() {
	right := n.getRight()

	left := n.getLeft()
	if left != nil {
		// sum := left.val + n.leftVal.val
		// left.val = sum
		// fmt.Print("left ", left.val, " ")
		left.add(n.leftVal.val)
	}
	if right != nil {
		// sum := right.val + n.rightVal.val
		// right.val = sum
		// fmt.Print("right ", right.val)
		right.add(n.rightVal.val)
	}
	// fmt.Println()

	deleteLeft := false
	if n.parent.leftNum != nil && n.parent.leftNum == n {
		// p1 := n.parent.leftNum
		// p2 := n
		// fmt.Println(*n.parent.leftNum, *&n)
		// fmt.Println(p1, p2)
		deleteLeft = true
	}
	deleteRight := false
	if n.parent.rightNum != nil && n.parent.rightNum == n {
		// p1 := n.parent.rightNum
		// p2 := n
		// fmt.Println(*n.parent.rightNum, *&n)
		// fmt.Println(p1, p2)
		deleteRight = true
	}

	// if deleteLeft && deleteRight {
	// 	// fmt.Println("Trying to delete both")
	// 	fmt.Println(&n.parent.leftNum)
	// 	fmt.Println(&n)
	// 	fmt.Println(*(&n))
	// 	fmt.Println(n)
	// 	fmt.Println(*&n)

	// 	fmt.Println(&n.parent.rightNum == &n)
	// }

	if deleteLeft {
		n.parent.leftNum = nil
	}
	if deleteRight {
		n.parent.rightNum = nil
	}
}

func (n number) getMag() int {
	leftMag := 0
	rightMag := 0
	if n.leftNum != nil {
		leftMag = n.leftNum.getMag()
	} else {
		leftMag = n.leftVal.val
	}

	if n.rightNum != nil {
		rightMag = n.rightNum.getMag()
	} else {
		rightMag = n.rightVal.val
	}

	return leftMag*3 + rightMag*2
}

func (n number) deepCopy() *number{
	newN := newNum()
	if n.leftNum != nil {
		newLeft := n.leftNum.deepCopy()
		newLeft.parent = newN
		newN.leftNum = newLeft
	} else {
		newN.leftVal.val = n.leftVal.val
	}

	if n.rightNum != nil {
		newRight := n.rightNum.deepCopy()
		newRight.parent = newN
		newN.rightNum = newRight
	} else {
		newN.rightVal.val = n.rightVal.val
	}
	return newN
}

func parseInput18(f string) []*number {
	file, err := os.Open("./data/day18_"+ f +".txt")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// for each opened bracket, make a new number,
	// 		if next is number, make that left val
	// 		if next is bracket, make a new number, setting it as left of curren, and set new as current number
	// 		if next is comma number,make that the right val, set current number as parent
	// 		if next is comma bracket, make a new number, setting it as the right of current, and set new as current number
	// 		if next is closing bracket, set parent of current as current

	snailNums := make([]*number,0)
	var root *number = nil
	var curNum *number = nil
	const left = 0
	const right = 1
	side := left //

	for scanner.Scan() {
		next := scanner.Text()
		for _, v := range next{
			switch v {
			case '[':
				temp := newNum()
				if root == nil {root = temp}
				if curNum == nil {
					curNum = temp
				} else {
					temp.parent = curNum
					if side == left{
						curNum.leftNum = temp
					} else {
						curNum.rightNum = temp
					}
					curNum = temp
				}

				side = left
			case ']':
				temp := curNum
				if curNum.parent != nil {
					curNum = curNum.parent
				} else { break }
				if temp == curNum.leftNum{
					side = left
				} else {
					side = right
				}
			case ',':
				side = right
			default:
				if side == left {
					val, _ := strconv.Atoi(string(v))
					curNum.leftVal = numVal{val}
				} else {
					val, _ := strconv.Atoi(string(v))
					curNum.rightVal = numVal{val}
				}
			}
		}
		snailNums = append(snailNums, root)
		root = nil
		curNum = nil
	}
	return snailNums
}
