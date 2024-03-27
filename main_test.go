package main

import (
	"testing"
)

func setupBoard(moves []struct{row, col int; player string}) {
	initializeBoard()
	for _, move := range moves {
		board[move.row][move.col] = move.player
	}
}

func TestIsGameOver(t *testing.T) {
	setupBoard([]struct{row, col int; player string}{
		{0, 0, "X"}, {0, 1, "X"}, {0, 2, "X"},
	})

	player = "X"
	if !isGameOver() {
		t.Errorf("Expected game to be over.")
	}
}

func TestSwitchPlayer(t *testing.T) {
	player = "X"
	switchPlayer()
	if player != "O" {
		t.Errorf("Expected player to be \"O\".")
	}

	switchPlayer()
	if player != "X" {
		t.Errorf("Expected player to be \"X\".")
	}
}
