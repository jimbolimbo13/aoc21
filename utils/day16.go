package utils

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"os"

	bitio "github.com/icza/bitio"
)

func Day16() {
	// d16_part1()
	d16_part2()
}

type state struct {
	f func ()
}

type subPacket struct {
	opType int
	remaining int
	retIx int
	values []int
	opID int
	sP_ID int
}

type literalCount struct {
	written *bytes.Buffer
	writer *bitio.Writer
	open bool
	count int
}

const (
	lengthOp = 0
	packetOp = 1
)

var (
	subPackets []subPacket
	curSubIx int
	versions []int
	bitIx int
	litCount literalCount
	reader *bitio.Reader
	curState int
	lastType int
	sP_IDs int
)

func performOperation(values []int, opID int) int {
	result := 0

	switch opID {
	case 0:
		for _, v := range values {
			result += v
		}
	case 1:
		result = 1
		for _, v := range values {
			result = result * v
		}
	case 2:
		result = -1
		for _, v := range values {
			if result == -1 { result = v }
			if v < result { result = v }
		}
	case 3:
		result = -1
		for _, v := range values {
			if result == -1 { result = v }
			if v > result { result = v }
		}
	case 5:
		if values[0] > values[1] { result = 1 }
	case 6:
		if values[0] < values[1] { result = 1 }
	case 7:
		if values[0] == values[1] { result = 1 }
	}

	return result
}

func decrementLengths(d int) {
	nextIx := curSubIx
	for {
		if nextIx == -1 { break }
		if subPackets[nextIx].opType == lengthOp {
			subPackets[nextIx].remaining -= d
		}

		nextIx = subPackets[nextIx].retIx
	}

	// for v := range subPackets{
	// 	if subPackets[v].opType == lengthOp && subPackets[v].remaining >= 0 {
	// 		subPackets[v].remaining -= d
	// 	}
	// }
	// if subPackets[curSubIx].opType == lengthOp && subPackets[curSubIx].remaining >= 0 {
	// 	subPackets[curSubIx].remaining -= d
	// }
}

func d16_part2(){
	reader = parseInput16("input")

	stateMachine := make(map[int]state)
	readVersion := func () {
		fmt.Println("remaining", subPackets[curSubIx].remaining, len(subPackets), subPackets[curSubIx].sP_ID)
		if curState > 0 && subPackets[curSubIx].retIx == -1 {
			// move to Done
			curState = -1
			return
		} else if subPackets[curSubIx].remaining <= 0 {
			// calculate, then remove
			fmt.Println("evaluating", subPackets[curSubIx])
			new_val := performOperation(subPackets[curSubIx].values, subPackets[curSubIx].opID)
			fmt.Println("got", new_val)
		 	// remove the subpacked
			// oldIx := curSubIx
			// fmt.Println("Removing subpackets", subPackets)
			// subPackets = subPackets[:oldIx]
			// fmt.Println( subPackets)
			curSubIx = subPackets[curSubIx].retIx
			// curSubIx = oldIx -1
			subPackets[curSubIx].values = append(subPackets[curSubIx].values, new_val)

			// for i := curSubIx; i >0; i--{
			for {
				// subPackets[curSubIx].remaining -= 1
				if subPackets[curSubIx].remaining <= 0 {
					fmt.Println("evaluating dead", subPackets[curSubIx])
					new_val := performOperation(subPackets[curSubIx].values, subPackets[curSubIx].opID)
					fmt.Println("got", new_val)
					curSubIx = subPackets[curSubIx].retIx
					subPackets[curSubIx].values = append(subPackets[curSubIx].values, new_val)
					// subPackets = subPackets[:i]
				} else {
					break
				}
			}

			if subPackets[curSubIx].retIx == -1 {
				// move to Done
				curState = -1
				return
			}
		}

		decode, _ := reader.ReadBits(3)
		decrementLengths(3)
		if subPackets[curSubIx].opType == packetOp{
			subPackets[curSubIx].remaining -= 1
		}
		versions = append(versions, int(decode))
		curState = 2
	}
	stateMachine[0] = state{readVersion}
	stateMachine[1] = state{readVersion}

	readType := func () {
		decode, _ := reader.ReadBits(3)
		lastType = int(decode)
		decrementLengths(3)
		if decode == 4 {
			litCount.count = 0
			curState = 3
		} else {
			curState = 4
		}
	}
	stateMachine[2] = state{readType}

	readLiteral := func () {
		more, _ := reader.ReadBits(1)
		decode, _ := reader.ReadBits(4)
		decrementLengths(5)
		if !litCount.open {
			litCount.written = &bytes.Buffer{}
			litCount.writer = bitio.NewWriter(litCount.written)
			litCount.open = true
		}

		litCount.writer.WriteBits(decode, 4)

		litCount.count ++
		if more == 1 {
			curState = 3
		} else {
			litCount.open = false
			litCount.writer.Close()
			r := bitio.NewReader(litCount.written)
			litTotal, _ := r.ReadBits(uint8(litCount.count*4))

			subPackets[curSubIx].values = append(subPackets[curSubIx].values, int(litTotal))

			fmt.Println("literal:", litTotal, litCount.count)
			curState = 1

		}
	}
	stateMachine[3] = state{readLiteral}

	readOperator := func() {
		decode, _ := reader.ReadBits(1)
		decrementLengths(1)
		if decode == 0 {
			curState = 5
		} else {
			curState = 6
		}
	}
	stateMachine[4] = state{readOperator}

	readLengthOp := func() {
		decode, _ := reader.ReadBits(15)
		decrementLengths(15)
		nextSubPacket := subPacket{
			opType:    lengthOp,
			remaining: int(decode),
			retIx:     curSubIx,
			values:    []int{},
			opID:      lastType,
			sP_ID:     sP_IDs,
		}
		sP_IDs ++
		fmt.Println("new subPacket", nextSubPacket)
		subPackets = append(subPackets, nextSubPacket)
		curSubIx = len(subPackets) - 1
		curState = 1
	}
	stateMachine[5] = state{readLengthOp}

	readPacketsOp := func() {
		decode, _ := reader.ReadBits(11)
		decrementLengths(11)
		nextSubPacket := subPacket{
			opType:    packetOp,
			remaining: int(decode),
			retIx:     curSubIx,
			values:    []int{},
			opID:      lastType,
			sP_ID:     sP_IDs,
		}
		sP_IDs ++
		fmt.Println("new subPacket", nextSubPacket)
		subPackets = append(subPackets, nextSubPacket)
		curSubIx = len(subPackets) - 1
		curState = 1
	}
	stateMachine[6] = state{readPacketsOp}

	// initializing stuff
	litCount = literalCount{
		written:    &bytes.Buffer{},
		writer: &bitio.Writer{},
		open:   false,
		count:  0,
	}
	curSubIx = 0
	sP_IDs = 0
	subPackets = make([]subPacket, 0)
	origPacket := subPacket{
		opType:    packetOp,
		remaining: 2,
		retIx:     -1,
		values:    []int{},
		opID:      0,
		sP_ID:     sP_IDs,
	}
	sP_IDs ++
	subPackets = append(subPackets, origPacket)
	curState = 0

	counter := 0
	for {
		// fmt.Println("state is:", curState)
		runState := stateMachine[curState]
		runState.f()
		if curState == -1 {break}
		counter ++
		// if counter > 30 {break}
	}
	ver_total := 0
	for _, v := range versions{
		ver_total += v
	}
	fmt.Println("versions", ver_total)
	final := subPackets[curSubIx].values
	fmt.Println("Day16, Part2:", final[0])
}

func d16_part1() {
	reader = parseInput16("input")

	stateMachine := make(map[int]state)
	readVersion := func () {
		decode, _ := reader.ReadBits(3)
		if subPackets[curSubIx].opType == lengthOp{
			subPackets[curSubIx].remaining -= 3
		} else {
			subPackets[curSubIx].remaining -= 1
		}
		// fmt.Println("remaining", subPackets[curSubIx].remaining, subPackets)
		if curState > 0 && subPackets[curSubIx].retIx == -1 {
			// move to Done
			curState = -1
			return
		} else if subPackets[curSubIx].remaining < 0 {
		 	// remove the subpacked
			oldIx := curSubIx
			// fmt.Println("Removing subpackets", oldIx, len(subPackets))
			subPackets = subPackets[:oldIx]
			// fmt.Println( subPackets)
			// curSubIx = subPackets[curSubIx].retIx
			curSubIx = oldIx -1

			for i := curSubIx; i >0; i--{
				subPackets[curSubIx].remaining -= 1
				if subPackets[curSubIx].remaining < 0 {
					curSubIx --
					subPackets = subPackets[:i]
				} else {
					break
				}
			}

			if subPackets[curSubIx].retIx == -1 {
				// move to Done
				curState = -1
				return
			}
		}
		versions = append(versions, int(decode))
		curState = 2
	}
	stateMachine[0] = state{readVersion}
	stateMachine[1] = state{readVersion}

	readType := func () {
		decode, _ := reader.ReadBits(3)
		if subPackets[curSubIx].opType == lengthOp{
			subPackets[curSubIx].remaining -= 3
		}
		if decode == 4 {
			litCount.count = 0
			curState = 3
		} else {
			curState = 4
		}
	}
	stateMachine[2] = state{readType}

	readLiteral := func () {
		more, _ := reader.ReadBits(1)
		decode, _ := reader.ReadBits(4)
		if subPackets[curSubIx].opType == lengthOp{
			subPackets[curSubIx].remaining -= 5
		}
		if !litCount.open {
			litCount.written = &bytes.Buffer{}
			litCount.writer = bitio.NewWriter(litCount.written)
			litCount.open = true
		}

		litCount.writer.WriteBits(decode, 4)

		litCount.count ++
		if more == 1 {
			curState = 3
		} else {
			litCount.open = false
			litCount.writer.Close()
			// r := bitio.NewReader(litCount.written)
			// litTotal, _ := r.ReadBits(uint8(litCount.count*4))

			// fmt.Println("literal:", litTotal, litCount.count)
			curState = 1

			// if subPackets[curSubIx].opType == packetOp {
			// 	subPackets[curSubIx].remaining -= 1
			// }
		}
	}
	stateMachine[3] = state{readLiteral}

	readOperator := func() {
		decode, _ := reader.ReadBits(1)
		if subPackets[curSubIx].opType == lengthOp{
			subPackets[curSubIx].remaining -= 1
		}
		if decode == 0 {
			curState = 5
		} else {
			curState = 6
		}
	}
	stateMachine[4] = state{readOperator}

	readLengthOp := func() {
		decode, _ := reader.ReadBits(15)
		if subPackets[curSubIx].opType == lengthOp{
			subPackets[curSubIx].remaining -= 15
		}
		nextSubPacket := subPacket{
			opType:    lengthOp,
			remaining: int(decode),
			retIx:     curSubIx,
			values:    []int{},
			opID:      0,
		}
		subPackets = append(subPackets, nextSubPacket)
		curSubIx = len(subPackets) - 1
		curState = 1
	}
	stateMachine[5] = state{readLengthOp}

	readPacketsOp := func() {
		decode, _ := reader.ReadBits(11)
		if subPackets[curSubIx].opType == lengthOp{
			subPackets[curSubIx].remaining -= 11
		}
		nextSubPacket := subPacket{
			opType:    packetOp,
			remaining: int(decode),
			retIx:     curSubIx,
			values:    []int{},
			opID:      0,
		}
		subPackets = append(subPackets, nextSubPacket)
		curSubIx = len(subPackets) - 1
		curState = 1
	}
	stateMachine[6] = state{readPacketsOp}

	// initializing stuff
	litCount = literalCount{
		written:    &bytes.Buffer{},
		writer: &bitio.Writer{},
		open:   false,
		count:  0,
	}
	curSubIx = 0
	subPackets = make([]subPacket, 0)
	origPacket := subPacket{
		opType:    packetOp,
		remaining: 2,
		retIx:     -1,
		values:    []int{},
		opID:      0,
	}
	subPackets = append(subPackets, origPacket)
	curState = 0

	counter := 0
	for {
		// fmt.Println("state is:", curState)
		runState := stateMachine[curState]
		runState.f()
		if curState == -1 {break}
		counter ++
		// if counter > 30 {break}
	}

	// fmt.Println(versions)
	ver_total := 0
	for _, v := range versions{
		ver_total += v
	}
	fmt.Println("\nDay16, Part1:", ver_total)
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
	var inputHex []byte
	for scanner.Scan() {
		check := scanner.Text()
		if check [:1] == "#" { continue }
		input = scanner.Bytes()
		inputHex = make([]byte, hex.DecodedLen(len(input)))
		_, err := hex.Decode(inputHex, input)
		if err != nil{
			fmt.Println(err)
			panic(err)
		}
	}

	// fmt.Println(inputHex)

	reader := bitio.NewReader(bytes.NewBuffer(inputHex))
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
// literal count
// 		bytes.Buffer for the number
// 		count of iterations
//
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
