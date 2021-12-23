package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const(
	POS = 0
	PTS = 1
	boardSize = 10
)

var(
	dieCount = 0
	diracWinPts = 3
	pointCounts = make([]int,30)
)

func Day21() {
	// d21_part1()
	d21_part2()
}

func d21_part2() {
	players := parseInput21("input")
// for each board position: ( 10 total  )
// 		for each count position: ( up to 30 )
// 		 	loop through  [0 0 0 1 3 6 7 6 3 1]
// 		 	 	board position + loop : (new pos)
// 		 	 	new points = new pos + count position ix
// 		 	 	if new points >= 21
// 		 	 	 	increment wins for that player
// 		 	 	else:
// 		 	 		increment board position [new pos] count position with new points*number in count position

	// create the board
	// for each player, for each board position, list of point values

	// we assume 2 players for now

	playerWins := make([]int,2)
	board := make([][][]int,2)
	for i := range board{
		board[i] = make([][]int,10)
		for j := range board[i] {
			board[i][j] = make([]int, 30)
		}
	}

	// fmt.Println(players)

	board[0][players[0][0]][players[0][1]] = 1
	board[1][players[1][0]][players[1][1]] = 1
	rolls := []int{0,0,0,1,3,6,7,6,3,1}
	cp := 0 // current player

	var offPlayer int
	// game loop
	for {
		cp = 0
		board, playerWins, offPlayer = diracTurn(board, rolls, cp, playerWins)
		// another round
		cp = 1
		board, playerWins, offPlayer = diracTurn(board, rolls, cp, playerWins)
		if offPlayer == 0 { break }
	}

	fmt.Println("\nWins:", playerWins)
	final := playerWins[0]
	if playerWins[0] < playerWins[1]{ final = playerWins[1] }

	fmt.Println("Day2, Part2:", final)
}

func diracTurn ( board [][][]int, rolls []int, cp int, playerWins []int ) ([][][]int, []int, int) {
	final := 1
	var newPos int
	var newPoints int
	var offPlayer float64 // how many rounds happened for other player
	newBoard := make([][][]int,2)
	for i := range newBoard{
		newBoard[i] = make([][]int,10)
		for j := range newBoard[i] {
			newBoard[i][j] = make([]int, 30)
		}
	}

	board1initial := 0
	board1total := 0
	board2total := 0

	for pos, positions := range board[cp] {
		for pts, points := range positions {
			if points > 0 {
				board1initial += (board[cp][pos][pts] )
				for r:=3; r<len(rolls); r++{
					newPos = boardPos(pos + r) // becomes second index in new board
					newPoints = addPoints(pts, newPos) // becomes the final index in new board
					if newPoints >= 21 {
						playerWins[cp] += (points*rolls[r])
					} else {
						newBoard[cp][newPos][newPoints] += (points*rolls[r])
						board1total += (points*rolls[r])
					}
					board[cp][pos][pts] = 0
				}
			}
		}
	}
	// fmt.Println("\nMultiplier values:", board1total, board1initial)
	if board1initial != 0 {
		offPlayer = (float64(board1total)/float64(board1initial))
		// fmt.Println(offPlayer, (float64(board1total)/float64(board1initial)))
	} else {
		offPlayer = 1
		final = 0
	}
	// fmt.Println(cp, (cp+1)%2)
	for pos, positions := range board[(cp+1)%2] {
		for pts, points := range positions {
			if points > 0 {
				// TODO multiply the number of points by how many offPlayer
				newBoard[(cp+1)%2][pos][pts] = int( float64(board[(cp+1)%2][pos][pts]) * offPlayer)
				board2total += int( float64(board[(cp+1)%2][pos][pts]) * offPlayer)
			}
		}
	}
	// for _, positions := range newBoard[cp] {
	// 	fmt.Println(positions)
	// }
	// fmt.Println("")
	// for _, positions := range newBoard[(cp+1)%2] {
	// 	fmt.Println(positions)
	// }
	// fmt.Println("")
	// fmt.Println("offPlayer:", offPlayer, "Board totals:", board1total, board2total, "\n")
	// fmt.Println(playerWins, )

	return newBoard, playerWins, final
}

func d21_part2_rec(){
	players := parseInput21("example")


	lazy := make(chan int)

	playerWins := make([]int, len(players))
	// playerWins := diracTurn(players, 0, 0, 0, 0)
	for i:=1; i<=3; i++{
		newWins := make([]int, len(players))
		newPlayers := deepCopyInt2(players)
		go diracTurnOld(newPlayers, i, 1, 0, 0, lazy, &newWins)
		data := <- lazy
		for i, v := range newWins {
			playerWins[i] += v
		}
		if data != -1 {

		}
	}

	fmt.Println(playerWins)
	fmt.Println(pointCounts, dieCount)
	fmt.Println("Day21, Part2:")
}

func d21_part1() {
	die := 0
	players := parseInput21("input")

	curPlayer := 0
	for {
		players[curPlayer] = runTurn(players[curPlayer], &die)
		if players[curPlayer][PTS] >= 1000 { break }
		curPlayer = (curPlayer + 1) % len(players)
	}

	fmt.Println(curPlayer, players)

	total := players[(curPlayer +1)%2][PTS] * dieCount
	fmt.Println("Day21, Part1:", total)
}

func diracTurnOld ( players [][]int, rollVal int, rollCount int, curPlayer int, points int, unlazy chan int, playerWins *[]int ) {
	// playerWins := make([]int, len(players))

	// fmt.Print("new turn:", players,rollVal, rollCount, curPlayer, points)
	if rollCount >= 3 {
		players[curPlayer][POS] = boardPos(players[curPlayer][POS] + rollVal)
		players[curPlayer][PTS] += addPoints(players[curPlayer][PTS], players[curPlayer][POS])
		pointCounts[rollVal] ++
		// fmt.Print("\n",players, playerWins)
		if players[curPlayer][PTS] >= diracWinPts {
			(*playerWins)[curPlayer] ++
			// fmt.Println("Is a winner!")
			unlazy <- 1
			return
		}
		rollVal = 0
		rollCount = 0
		curPlayer = (curPlayer + 1) % len(players)
		points = 0
		unlazy <- 1
		return
	} else {
		// fmt.Println("")
		points += rollVal
	}

	for i:=1; i<=3; i++{
		lazy := make(chan int)
		newWins := make([]int, len(players))
		newPlayers := deepCopyInt2(players)
		go diracTurnOld(newPlayers, rollVal + i, rollCount + 1, curPlayer, points, lazy, &newWins)
		data := <- lazy
		for i, v := range newWins {
			(*playerWins)[i] += v
		}
		if data == 1 {continue}
	}
	unlazy <- 1
	return
}

func deepCopyInt2( p [][]int ) [][]int {
	newP := make([][]int, len(p))
	for i, v := range p {
		for _, k := range v{
			newP[i] = append(newP[i], k)
		}
	}
	return newP
}

func runTurn(player []int, die *int ) []int {
	numRolls := 3
	rollVal := 0

	for i:=0; i<numRolls; i++{
		*die = dieRoll(*die)
		rollVal += *die
	}
	player[POS] = boardPos(player[POS] + rollVal)
	player[PTS] = addPoints(player[PTS], player[POS])
	return player
}

func dieRoll(d int) int {
	// deterministic die
	dieCount ++
	d ++
	if d > 100 {
		return d-100
	} else {
		return d
	}
}

func parseInput21(f string) [][]int {
	file, err := os.Open("./data/day21_"+ f +".txt")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	players := make([][]int, 0)
	for scanner.Scan() {
		next := make([]int,2)
		new := strings.Fields(scanner.Text())
		// player, _ := strconv.Atoi(new[1])
		pos, _ := strconv.Atoi(new[4])
		next[POS] = boardPos(pos)
		players = append(players, next)
	}

	return players
}

func addPoints ( pts int, more int ) int {
	if more == 0 {
		return pts + boardSize
	} else {
		return pts + more
	}
}

func boardPos ( p int ) int {
	return p - (p/boardSize)*boardSize
}



// 1 2 3
//
// 1						2						3
// 2 3 4					3 4 5					4 5 6
//
// 2		3		4		3		4		5		4		5		6
// 3 4 5	4 5 6	5 6 7	4 5 6	5 6 7	6 7 8	5 6 7	6 7 8	7 8 9
//
// 27 possibilities per turn

// array for each player's positions, for each position, an array of points
// all possible rolls
// [0 0 0 1 3 6 7 6 3 1]
//
// for each position, all possible added points
//
// 27 different outcomes, but not all unique
//
// 7 outcomes involve rolling 6 total, moving 6 spaces, then adding the score
// in the end, there are 7 new possible board positions
//
// the next person goes and there are now 27 more instances for each position on the board
//
// p1 turn again, there are 27^27 instances
//
// for each instance there are 27 outcomes, but not all unique
// 		for each postion, there are 27 new outcomes
// 		 	2 on position 1 with 4 points
// 		 	3 on position 1 with 2 points:
// 		 	 	there are now 3 on 4 with 6 points for each
// 		 	 	9 on 5 with 7 points each
// 		 	 	18 on 6 with 8 points each
// 		 	 	 etc
// 		 	 plus
// 		 	 	2 more on 4 with 8 points
// 		 	 	6 more on 5 with 9 points
// 		 	 	12 more on 6 with 10 points
//
// for each board position: ( 10 total  )
// 		for each count position: ( up to 30 )
// 		 	loop through  [0 0 0 1 3 6 7 6 3 1]
// 		 	 	board position + loop : (new pos)
// 		 	 	new points = new pos + count position ix
// 		 	 	if new points >= 21
// 		 	 	 	increment wins for that player
// 		 	 	else:
// 		 	 		increment board position [new pos] count position with new points*number in count position
