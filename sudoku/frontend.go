package main

import (
	"fmt"
	"strconv"
	"syscall/js"
	"time"
)

//Checks to see if any entered fields could be true.
//If they could be true with the given clues, they're
//ignored. If they can't be true, they're removed.
func checkJS(x js.Value, y []js.Value) interface{} {
	boardSize := len(gameBoard)
	var fields []Position

	for i := 0; i < boardSize; i++ {
		for n := 0; n < boardSize; n++ {
			if gameBoard[i][n] == 0 {
				fields = append(fields, Position{X: i, Y: n})
			}
		}
	}

	for i := 0; i < len(fields); i++ {
		field := fields[i]

		cellIndex := strconv.Itoa(field.X*len(gameBoard) + field.Y)

		if js.Global().Get("document").Call("getElementById", "cell-"+cellIndex).Get("value").Truthy() {
			value, err := strconv.Atoi(js.Global().Get("document").Call("getElementById", "cell-"+cellIndex).Get("value").String())
			if err != nil {
				fmt.Println("The value entered into cell-" + cellIndex + " was not valid.")
				js.Global().Get("document").Call("getElementById", "cell-"+cellIndex).Set("value", js.ValueOf(nil))
			} else {
				if !validEntry(gameBoard, value, Position{X: field.X, Y: field.Y}) {
					js.Global().Get("document").Call("getElementById", "cell-"+cellIndex).Set("value", js.ValueOf(nil))
				}
			}
		}
	}

	return nil
}

//Bruteforces the sudoku puzzle, and displays each step
func solveJS(x js.Value, y []js.Value) interface{} {
	boardSize := len(gameBoard)

	outputGameBoard()

	for i := 0; i < boardSize; i++ {
		for n := 0; n < boardSize; n++ {
			cellIndex := strconv.Itoa(i*boardSize + n)
			js.Global().Get("document").Call("getElementById", "cell-"+cellIndex).Set("disabled", js.ValueOf(true))
		}
	}

	go SolveSodoku(gameBoard, (time.Millisecond * 100))

	return nil
}

//Generates and displays a new board
func newJS(x js.Value, y []js.Value) interface{} {
	newPuzzle()

	return nil
}

//Generates and displays a new board
func newPuzzle() {
	solutionBoard, gameBoard = GenerateSudoku(9, 50)

	outputGameBoard()
}

//Displays the board stored in
//the global "gameBoard" var
func outputGameBoard() {
	boardSize := len(gameBoard)

	for i := 0; i < boardSize; i++ {
		for n := 0; n < boardSize; n++ {
			cellIndex := strconv.Itoa(i*boardSize + n)

			if gameBoard[i][n] == 0 {
				js.Global().Get("document").Call("getElementById", "cell-"+cellIndex).Set("disabled", js.ValueOf(false))
				js.Global().Get("document").Call("getElementById", "cell-"+cellIndex).Set("value", js.ValueOf(nil))
			} else {
				js.Global().Get("document").Call("getElementById", "cell-"+cellIndex).Set("disabled", js.ValueOf(true))
				js.Global().Get("document").Call("getElementById", "cell-"+cellIndex).Set("value", js.ValueOf(gameBoard[i][n]))
			}
		}
	}
}
