package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	bitio "github.com/icza/bitio"
)

var(
	infinite int
)
func Day20() {
	d20_part1()
	d20_part2()
}

func d20_part2(){
	infinite = 0
	image, enhancement := parseInput20("input")
	enhRuns := 50

	image = enhanceImage(padImage(image, 3), enhancement)

	for i:=1; i<enhRuns; i++{
		image = enhanceImage(padImage(image, 1), enhancement)
	}

	total := 0
	for _, v := range image {
		for _, k := range v {
			total += k
		}
	}
	fmt.Println("Day20, Part2:", total)
}

func d20_part1() {
	infinite = 0
	image, enhancement := parseInput20("input")

	enhImg := enhanceImage(padImage(image, 3), enhancement)

	enhImg = enhanceImage(padImage(enhImg, 1), enhancement)

	total := 0
	for _, v := range enhImg {
		for _, k := range v {
			total += k
		}
	}

	fmt.Println("Day20, Part1:", total)
}

func enhanceImage(img [][]int, enh []int) [][]int {
	enhImg := make([][]int,len(img))
	for i, v := range img {
		enhImg[i] = make([]int, len(v))
		for j := range v{
			enhPx := decodPixel(img, i, j, enh)
			enhImg[i][j] = enhPx
		}
	}
	infinite = enhImg[0][0]
	return enhImg
}

func decodPixel(img [][]int, y int, x int, enh []int) int {

	// 1 2 3
	// 4 5 6
	// 7 8 9
	//
	//5-row+1  5-row  5-row-1
	//5-1       5+1
	//5+row-1  5+row  5+row+1

	b := bytes.Buffer{}
	newRow := bitio.NewWriter(&b)

	newX := 0
	newY := 0
	dim := []int{len(img[0]), len(img)}
	for i := -1; i <= 1; i++{ // row diff
		for j := -1; j <= 1; j++{ // col diff
			newX = x + j
			newY = y + i
			if IsValidIxGrid(newX, newY, dim){
				newRow.WriteBits(uint64(img[newY][newX]), 1)
			} else {
				newRow.WriteBits(uint64(infinite), 1)
			}

		}

	}
	newRow.Close()
	enhIx, _ := bitio.NewReader(&b).ReadBits(9)
	return int(enh[enhIx])
}

func padImage(img [][]int, padding int) ([][]int) {
	padded := make([][]int, len(img) + 2*padding)

	for i := range padded{
		newRow := make([]int, len(img[0]) + 2*padding)

		for j := range newRow {
			if (padding) <= i && i < len(padded) - (padding){
				if (padding) <= j && j < len(padded[0]) - (padding){
					newRow[j] = img[i-padding][j-padding]
				} else {
					newRow[j] = infinite
				}
			} else {
				newRow[j] = infinite
			}
		}

		padded[i] = newRow
	}
	return padded
}

func parseInput20(f string) ([][]int, []int) {
	file, err := os.Open("./data/day20_"+ f +".txt")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	first_line := strings.Split(scanner.Text(), "")

	enhancement := make([]int,len(first_line))
	for i, v := range first_line{
		enhancement[i] = decodeInput(v)
	}

	// skip the blank line
	scanner.Scan()

	image := make([][]int, 0)
	for scanner.Scan() {
		next_row := strings.Split(scanner.Text(),"")
		img_row := make([]int, len(next_row))
		for i, v := range next_row {
			img_row[i] = decodeInput(v)
		}
		image = append(image, img_row)
	}

	return image, enhancement
}

func decodeInput ( v string ) int {
		switch v {
		case "#":
			return 1
		case ".":
			return 0
		default:
			panic("Bad string input in decodeInput")
		}
}

func encodeInput ( v int ) string {
	if v == 1 {
		return "#"
	} else if v == 0 {
		return "."
	} else {
		panic("Bad int input to encodeInput")
	}
}

func dumpImage (img [][]int) {
	for _, v := range img {
		for _, c := range v {
			fmt.Print(encodeInput(c))
		}
		fmt.Println()
	}
}
