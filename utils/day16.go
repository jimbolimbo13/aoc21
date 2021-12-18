package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"encoding/hex"
	bitio "github.com/icza/bitio"
)

func Day16() {
	d16_part1()
	d16_part2()
}

func d16_part2(){
	parseInput16("example")
	fmt.Println("Day16, Part2:")
}

func d16_part1() {
	reader := parseInput16("example")
	next, err := reader.ReadBits(3)
	if err != nil{
		panic(err)
	}
	fmt.Print(next, " ")

	next, err = reader.ReadBits(3)
	if err != nil{
		panic(err)
	}
	fmt.Print(next, " ")
	next, err = reader.ReadBits(1)
	if err != nil{
		panic(err)
	}
	fmt.Print(next, " ")
	next, err = reader.ReadBits(4)
	if err != nil{
		panic(err)
	}
	fmt.Print(next, " ")
	next, err = reader.ReadBits(1)
	if err != nil{
		panic(err)
	}
	fmt.Print(next, " ")
	next, err = reader.ReadBits(4)
	if err != nil{
		panic(err)
	}
	fmt.Print(next, " ")
	next, err = reader.ReadBits(1)
	if err != nil{
		panic(err)
	}
	fmt.Print(next, " ")
	next, err = reader.ReadBits(4)
	if err != nil{
		panic(err)
	}
	fmt.Print(next, " ")

	fmt.Println("\nDay16, Part1:")
}

func parseInput16(f string) *bitio.Reader {
	file, err := os.Open("./data/day16_"+ f +".txt")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []byte
	for scanner.Scan() {
		check := scanner.Text()
		if check [:1] == "#" { continue }
		input = scanner.Bytes()
	}

	fmt.Println(input)

	reader := bitio.NewReader(bytes.NewBuffer(input))
	return reader
}

// State machine
//
// keeping track
// sub-packets []
// 		list where each position is for a new set of sub-packets
// 		basically a sub-packet stack
// 		needs to contain struct with:
//			operator type
//			remaining count
//			return ix
// 		gets initialized before starting
//
// current sub-packet ix
//
// versions []
// 		list of all the version numbers encountered
//
// bit index - int
// 		keeps track of our current index
//
// 1 readVersion
// 		if sub-packets is empty, move to Done
// 		read 3 bits, record version number in versions
//		if sub-packets.type is Length, decrement by 3
//		move to readType
// 2 readType
// 		read 3 bits, decode packet type
//		if sub-packets.type is Length, decrement by 3
//		else decrement by 1
// 		if 4, initialize literal value, move to literal
// 		else, operator
// 3 literal
// 		read 5 bits, note first bit and last 4
//		if sub-packets.type is Length, decrement by 5
// 		append 4 bits to literal value
// 		if first bit is 1, move to literal
// 		if first bit is 0:
// 			decode the literal value
// 			move to readVersion
// 4 readOperator
// 		read 1 bit
//		if sub-packets.type is Length, decrement by 4
// 		if 0, move to readLengthOp
// 		if 0, move to readPacketsOp
// 5 readLenthOp
// 		read 15 bits, decode number,
//		if sub-packets.type is Length, decrement by 15
// 		append to sub-packets:
// 			{lengthOp,
// 			length (from decoded number),
// 			return ix (current sub-packet ix)}
// 		set sub-packet ix
// 		move to readVersion
// 6 readPacketsOp
// 		read 11 bits, decode number,
//		if sub-packets.type is Length, decrement by 11
// 		append to sub-packets:
// 			{packetsOp,
// 			length (from decoded number),
// 			return ix (current sub-packet ix)}
// 		set sub-packet ix
// 		move to readVersion
//
// functions
// 1 read - takes number of bits, moves the bit index, returns bits read
