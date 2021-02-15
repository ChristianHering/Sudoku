package sudoku

import (
	"math/rand"
)

//GameBoards holds a solution board in the first index
//and the unfinished game board in the second index
type GameBoards [2]Board

//Board holds the values of each position of a 9x9 sodoku board
type Board [][]int

//PositionValue holds a board cell position and it's value
type PositionValue struct {
	Value int
	Pos   Position
}

//Position stores a single cell from a sudoku board
type Position struct {
	X int
	Y int
}

//NewGame returns 2 Sudoku Boards
//
//The first board is the game's solution board while
//the second board is the non complete game board
func NewGame() GameBoards {
	return GameBoards{}
}

//GenerateSudoku a new random sodoku board, and return both the filled out solution board
//and the game board.
//
//TODO: Ensure the game board returned only has 1 solution
func (gameBoards *GameBoards) GenerateSudoku(boardSize int, clues int) {
	gameBoards[0] = Board{
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

	cells, _ := gameBoards[0].NewBoard()

	gameBoards[1] = gameBoards[0]

	for i := 0; i < (boardSize*boardSize)-clues; i++ { //Remove random values until you have X clues left
		cellIndex := rand.Intn(len(cells))

		cell := cells[cellIndex]

		cells[cellIndex] = cells[len(cells)-1]
		cells = cells[:len(cells)-1] //Remove the value at cell index from our cell slice

		gameBoards[1][cell.X][cell.Y] = 0
	}
}

//NewBoard generates a new sudoku board that's 100% filled in
//
//All values in the board's int array must be 0
func (board *Board) NewBoard() (filledCells []Position, filled bool) {
	field := board.nextEmptyField()
	if (field == Position{X: -1, Y: -1}) {
		return nil, true //This is called when the puzzle is filled
	}

	var validValues []int

	for i := 0; i < len(*board); i++ {
		if board.ValidEntry(i+1, Position{X: field.X, Y: field.Y}) {
			validValues = append(validValues, i+1)
		}
	}

	rand.Shuffle(len(validValues), func(i, n int) { validValues[i], validValues[n] = validValues[n], validValues[i] })

	for i := 0; i < len(validValues); i++ {
		(*board)[field.X][field.Y] = validValues[i]

		filledCells, filled := board.NewBoard()
		if filled == true {
			filledCells = append(filledCells, field)

			return filledCells, true
		}

		(*board)[field.X][field.Y] = 0
	}

	return nil, false
}

//SolveSudoku recursivly loops through the game board,
//bruteforcing each position until it's solved
func (board *Board) SolveSudoku(solutionPath *[]PositionValue) (*[]PositionValue, bool) {
	field := board.nextEmptyField()
	if (field == Position{X: -1, Y: -1}) {
		return solutionPath, true //This is called when the puzzle is solved
	}

	for i := 0; i < len(*board); i++ {
		if board.ValidEntry(i+1, Position{X: field.X, Y: field.Y}) {
			(*board)[field.X][field.Y] = i + 1

			*solutionPath = append(*solutionPath, PositionValue{Value: i + 1, Pos: field})

			solutionPath, solved := board.SolveSudoku(solutionPath)
			if solved == true {
				return solutionPath, true
			}

			(*board)[field.X][field.Y] = 0

			*solutionPath = append(*solutionPath, PositionValue{Value: 0, Pos: field})
		}
	}

	return solutionPath, false
}

//ValidEntry checks if "value" is a valid entry for the
//position "position" given a sudoku board
func (board *Board) ValidEntry(value int, position Position) bool {
	length := len(*board)

	for i := 0; i < length; i++ { //Y axis check
		if (*board)[position.X][i] == value && position.Y != i {
			return false
		}
	}

	for i := 0; i < length; i++ { //X axis check
		if (*board)[i][position.Y] == value && position.X != i {
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
			if (*board)[i][n] == value && (Position{X: i, Y: n} != position) {
				return false
			}
		}
	}

	return true
}

//nextEmptyField loops through every item, and returns the first
//index position that holds a value of 0
//
//TODO: Currently using this for a board has a big O of n^2, room for improvement
func (board *Board) nextEmptyField() Position {
	for i := 0; i < len(*board); i++ {
		for n := 0; n < len((*board)[0]); n++ {
			if (*board)[i][n] == 0 {
				return Position{X: i, Y: n}
			}
		}
	}

	return Position{X: -1, Y: -1} //-1 is returned when there are no more empty spaces, and the board is already solved
}
