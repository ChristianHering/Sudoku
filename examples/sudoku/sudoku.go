package main

import (
	"math/rand"
	"strconv"
	"syscall/js"
	"time"
)

type Position struct {
	X int
	Y int
}

//Solve recursivly loops through the game board,
//bruteforcing each position until it's solved
func SolveSodoku(board [][]int, stepInterval time.Duration) ([][]int, bool) {
	field := nextEmptyField(board)
	if (field == Position{X: -1, Y: -1}) {
		return board, true //This is called when the puzzle is solved
	}

	cellIndex := strconv.Itoa(field.X*len(board) + field.Y)

	for i := 0; i < len(board); i++ {
		if validEntry(board, i+1, Position{X: field.X, Y: field.Y}) {
			board[field.X][field.Y] = i + 1

			if stepInterval > (time.Nanosecond * 100) {
				js.Global().Get("document").Call("getElementById", "cell-"+cellIndex).Set("value", js.ValueOf(strconv.Itoa(i+1)))

				time.Sleep(stepInterval)
			}

			solvedBoard, solved := SolveSodoku(board, stepInterval)
			if solved == true {
				return solvedBoard, true
			}

			board[field.X][field.Y] = 0

			if stepInterval > (time.Nanosecond * 100) {
				js.Global().Get("document").Call("getElementById", "cell-"+cellIndex).Set("value", js.ValueOf(strconv.Itoa(0)))

				time.Sleep(stepInterval)
			}
		}
	}

	return nil, false
}

//Generate a new random sodoku board, and return both the filled out solution board
//and the game board. TODO make sure the game board returned only has 1 solution
func GenerateSudoku(boardSize int, clues int) (solutionBoard [][]int, gameBoard [][]int) {
	sBoard := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	sBoard, cells, _ := fillBoard(sBoard)

	gBoard := sBoard

	for i := 0; i < (boardSize*boardSize)-clues; i++ { //Remove random values until you have X clues left
		rand.Seed(time.Now().UnixNano())

		cellIndex := rand.Intn(len(cells))

		cell := cells[cellIndex]

		cells[cellIndex] = cells[len(cells)-1]
		cells = cells[:len(cells)-1] //Remove the value at cell index from our cell slice

		gBoard[cell.X][cell.Y] = 0
	}

	return sBoard, gBoard
}

//Generates a sudoku board that's 100% filled in
func fillBoard(board [][]int) (filledBoard [][]int, filledCells []Position, filled bool) {
	field := nextEmptyField(board)
	if (field == Position{X: -1, Y: -1}) {
		return board, nil, true //This is called when the puzzle is filled
	}

	var validValues []int

	for i := 0; i < len(board); i++ {
		if validEntry(board, i+1, Position{X: field.X, Y: field.Y}) {
			validValues = append(validValues, i+1)
		}
	}

	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(validValues), func(i, n int) { validValues[i], validValues[n] = validValues[n], validValues[i] })

	for i := 0; i < len(validValues); i++ {
		board[field.X][field.Y] = validValues[i]

		filledBoard, filledCells, filled := fillBoard(board)
		if filled == true {
			filledCells = append(filledCells, field)

			return filledBoard, filledCells, true
		}

		board[field.X][field.Y] = 0
	}

	return nil, nil, false
}

//Checks if "value" is a valid entry for the
//position "position" given a sudoku board
func validEntry(board [][]int, value int, position Position) bool {
	length := len(board)

	for i := 0; i < length; i++ { //Y axis check
		if board[position.X][i] == value && position.Y != i {
			return false
		}
	}

	for i := 0; i < length; i++ { //X axis check
		if board[i][position.Y] == value && position.X != i {
			return false
		}
	}

	nonetX := int(position.X / 3)
	nonetY := int(position.Y / 3)

	//This looks a bit funky, but we're essentially finding the index
	//position of the nonet that holds our position vector, then ranging
	//through that nonet's 3x3 grid to see if our value is in it already
	for i := nonetX * 3; i < nonetX*3+3; i++ {
		for n := nonetY * 3; n < nonetY*3+3; n++ {
			if board[i][n] == value && (Position{X: i, Y: n} != position) {
				return false
			}
		}
	}

	return true
}

//Loops through every item, and returns the first
//index position that holds a value of 0
//
//Currently this has a big O of n^2, room for improvement TODO
func nextEmptyField(board [][]int) Position {
	for i := 0; i < len(board); i++ {
		for n := 0; n < len(board[0]); n++ {
			if board[i][n] == 0 {
				return Position{X: i, Y: n}
			}
		}
	}

	return Position{X: -1, Y: -1} //-1 is returned when there are no more empty spaces, and the board is already solved
}
