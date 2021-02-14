package sudoku

import (
	"testing"
)

func TestNewGame(t *testing.T) {
	boards := NewGame()

	for boardIndex := 0; boardIndex < len(boards); boardIndex++ {
		for indexX := 0; indexX < len(boards); indexX++ {
			for indexY := 0; indexY < len(boards[boardIndex]); indexY++ {
				if int(boards[boardIndex][indexX][indexY]) != 0 {
					t.Errorf("NewGame() Returned %d instead of 0", boards[boardIndex][indexX][indexY])
				}
			}
		}
	}
}

func TestValidEntry(t *testing.T) {
	board := Board{ //Valid check
		{2, 8, 6, 3, 9, 4, 1, 7, 5},
		{3, 4, 5, 8, 1, 7, 2, 6, 9},
		{7, 1, 9, 5, 2, 6, 3, 4, 8},
		{5, 2, 7, 6, 8, 3, 9, 1, 4},
		{1, 3, 8, 9, 4, 5, 6, 2, 7},
		{9, 6, 4, 2, 7, 1, 5, 8, 3},
		{6, 9, 2, 4, 3, 8, 7, 5, 1},
		{8, 7, 3, 1, 5, 2, 4, 9, 6},
		{4, 5, 1, 7, 6, 9, 8, 3, 2},
	}

	for i := 0; i < len(board)*len(board); i++ {
		position := Position{X: i / len(board), Y: i % len(board)}

		if board.ValidEntry(board[i/len(board)][i%len(board)], position) != true {
			t.Error("ValidEntry() found an invalid entry in a valid board!")
		}
	}

	board = Board{ //Nonet test
		{1, 1, 1, 2, 2, 2, 3, 3, 3},
		{1, 1, 1, 2, 2, 2, 3, 3, 3},
		{1, 1, 1, 2, 2, 2, 3, 3, 3},
		{4, 4, 4, 5, 5, 5, 6, 6, 6},
		{4, 4, 4, 5, 5, 5, 6, 6, 6},
		{4, 4, 4, 5, 5, 5, 6, 6, 6},
		{7, 7, 7, 8, 8, 8, 9, 9, 9},
		{7, 7, 7, 8, 8, 8, 9, 9, 9},
		{7, 7, 7, 8, 8, 8, 9, 9, 9},
	}

	for i := 0; i < len(board)*len(board); i++ {
		position := Position{X: i / len(board), Y: i % len(board)}

		if board.ValidEntry(board[i/len(board)][i%len(board)], position) != false {
			t.Error("ValidEntry() didn't find an invalid cell in our nonet board!")
		}
	}

	board = Board{ //X axis test
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{2, 2, 2, 2, 2, 2, 2, 2, 2},
		{3, 3, 3, 3, 3, 3, 3, 3, 3},
		{4, 4, 4, 4, 4, 4, 4, 4, 4},
		{5, 5, 5, 5, 5, 5, 5, 5, 5},
		{6, 6, 6, 6, 6, 6, 6, 6, 6},
		{7, 7, 7, 7, 7, 7, 7, 7, 7},
		{8, 8, 8, 8, 8, 8, 8, 8, 8},
		{9, 9, 9, 9, 9, 9, 9, 9, 9},
	}

	for i := 0; i < len(board)*len(board); i++ {
		position := Position{X: i / len(board), Y: i % len(board)}

		if board.ValidEntry(board[i/len(board)][i%len(board)], position) != false {
			t.Error("ValidEntry() didn't find an invalid cell in our X axis board!")
		}
	}

	board = Board{ //Y axis test
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	for i := 0; i < len(board)*len(board); i++ {
		position := Position{X: i / len(board), Y: i % len(board)}

		if board.ValidEntry(board[i/len(board)][i%len(board)], position) != false {
			t.Error("ValidEntry() didn't find an invalid cell in our Y axis board!")
		}
	}
}

func TestNextEmptyField(t *testing.T) {
	board := Board{
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

	for i := 0; i < len(board)*len(board); i++ {
		position := board.nextEmptyField()

		if board[position.X][position.Y] != 0 {
			t.Errorf("NextEmptyField() didn't return the correct Position!\nReturned X: %d\nReturned Y: %d\nExpected X: 1\nExpected Y: 1\n", position.X, position.Y)
		}

		board[i/len(board)][i%len(board)] = 1
	}

	position := board.nextEmptyField()
	if position.X != -1 {
		t.Errorf("NextEmptyField() returned a non -1 value when all values should have been filled. Position returned: X-%d Y-%d", position.X, position.Y)
	}

	board[3][3] = 0

	position = board.nextEmptyField()

	if position.X != 3 || position.Y != 3 {
		t.Errorf("NextEmptyField() returned X- %d and Y- %d when it should have returned X- 3 Y- 3", position.X, position.Y)
	}
}
