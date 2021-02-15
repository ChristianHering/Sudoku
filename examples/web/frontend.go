package main

import (
	"fmt"
	"strconv"
	"syscall/js"
	"time"

	sudoku "github.com/ChristianHering/Sudoku"
)

//Checks to see if any entered fields could be true.
//If they could be true with the given clues, they're
//ignored. If they can't be true, they're removed.
func checkJS(x js.Value, y []js.Value) interface{} {
	boardSize := len(gameBoards[1])
	var fields []sudoku.Position

	for i := 0; i < boardSize; i++ {
		for n := 0; n < boardSize; n++ {
			if gameBoards[1][i][n] == 0 {
				fields = append(fields, sudoku.Position{X: i, Y: n})
			}
		}
	}

	for i := 0; i < len(fields); i++ {
		field := fields[i]

		cellIndex := strconv.Itoa(field.X*len(gameBoards[1][1]) + field.Y)

		if js.Global().Get("document").Call("getElementById", "cell-"+cellIndex).Get("value").Truthy() {
			value, err := strconv.Atoi(js.Global().Get("document").Call("getElementById", "cell-"+cellIndex).Get("value").String())
			if err != nil {
				fmt.Println("The value entered into cell-" + cellIndex + " was not valid.")
				js.Global().Get("document").Call("getElementById", "cell-"+cellIndex).Set("value", js.ValueOf(nil))
			} else {
				if !gameBoards[1].ValidEntry(value, sudoku.Position{X: field.X, Y: field.Y}) {
					js.Global().Get("document").Call("getElementById", "cell-"+cellIndex).Set("value", js.ValueOf(nil))
				}
			}
		}
	}

	return nil
}

//Bruteforces the sudoku puzzle, and displays each step
func solveJS(x js.Value, y []js.Value) interface{} {
	boardSize := len(gameBoards[1])

	outputGameBoard()

	for i := 0; i < boardSize; i++ {
		for n := 0; n < boardSize; n++ {
			cellIndex := strconv.Itoa(i*boardSize + n)
			js.Global().Get("document").Call("getElementById", "cell-"+cellIndex).Set("disabled", js.ValueOf(true))
		}
	}

	var moves []sudoku.PositionValue

	_, _ = gameBoards[1].SolveSudoku(&moves)
	fmt.Println(moves)
	go func() {
		for i := 0; i < len(moves); i++ {
			move := moves[i]
			fmt.Println(move)
			if move.Value != 0 {
				js.Global().Get("document").Call("getElementById", "cell-"+(strconv.Itoa((move.Pos.X*9)+move.Pos.Y))).Set("value", js.ValueOf(move.Value))
			} else {
				js.Global().Get("document").Call("getElementById", "cell-"+(strconv.Itoa((move.Pos.X*9)+move.Pos.Y))).Set("value", js.ValueOf("0"))
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

	return nil
}

//Generates and displays a new board
func newJS(x js.Value, y []js.Value) interface{} {
	newPuzzle()

	return nil
}

//Generates and displays a new board
func newPuzzle() {
	gameBoards = sudoku.NewGame()

	gameBoards.GenerateSudoku(9, 50)

	outputGameBoard()
}

//Displays the board stored in
//the global "gameBoard" var
func outputGameBoard() {
	boardSize := len(gameBoards[1])

	for i := 0; i < boardSize; i++ {
		for n := 0; n < boardSize; n++ {
			cellIndex := strconv.Itoa(i*boardSize + n)

			if gameBoards[1][i][n] == 0 {
				js.Global().Get("document").Call("getElementById", "cell-"+cellIndex).Set("disabled", js.ValueOf(false))
				js.Global().Get("document").Call("getElementById", "cell-"+cellIndex).Set("value", js.ValueOf(nil))
			} else {
				js.Global().Get("document").Call("getElementById", "cell-"+cellIndex).Set("disabled", js.ValueOf(true))
				js.Global().Get("document").Call("getElementById", "cell-"+cellIndex).Set("value", js.ValueOf(gameBoards[1][i][n]))
			}
		}
	}
}
