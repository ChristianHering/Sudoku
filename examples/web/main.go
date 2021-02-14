package main

import (
	"syscall/js"

	sudoku "github.com/ChristianHering/Sudoku"
)

var gameBoards sudoku.GameBoards

//In order to build this file directly for the webapp,
//run the following command in the root of this package:
//
//go build -o ./../src/asm/sudoku.wasm ./
func main() {
	newPuzzle()

	check := js.FuncOf(checkJS)
	solve := js.FuncOf(solveJS)
	new := js.FuncOf(newJS)

	js.Global().Get("document").Call("getElementById", "checkButton").Call("addEventListener", "click", check)
	js.Global().Get("document").Call("getElementById", "solveButton").Call("addEventListener", "click", solve)
	js.Global().Get("document").Call("getElementById", "newButton").Call("addEventListener", "click", new)

	<-make(chan struct{})
}
