package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	board [3][3]string
	player = "X"
)

func main() {
	initializeBoard()
	printBoard()

	fmt.Println("Enter row and column (0-2) separated by space")
	for {
		fmt.Printf("Player %s\n", player)
		row, col := getPlayerMove()
		if board[row][col] == "" {
			board[row][col] = player
			printBoard()
			if isGameOver() {
				fmt.Println("Player", player, "wins")
				return
			} else if isGameDraw() {
				fmt.Println("It's a draw")
				return
			}

			switchPlayer();
		} else {
			fmt.Printf("Cell (%d, %d) is occupied by %s\n", row, col, board[row][col])
		}
	}
}

func initializeBoard() {
	for i := range board {
		for j := range board[i] {
			board[i][j] = ""
		}
	}
}

func printBoard() {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == "" {
				fmt.Print("- ")
			} else {
				fmt.Print(board[i][j], " ")
			}
		}

		fmt.Println()
	}
}

func getPlayerMove() (int, int) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")
	row, _ := strconv.Atoi(parts[0])
	col, _ := strconv.Atoi(parts[1])
	return row, col
}

func switchPlayer() {
	if player == "X" {
		player = "O"
	} else {
		player = "X"
	}
}

func areEquals(str1 string, str2 string, str3 string) bool {
	return str1 == str2 && str2 == str3
}

func isGameOver() bool {
	for i:= range board {
		isHorizontalWin := board[0][i] == player && areEquals(board[0][i], board[1][i], board[2][i])
		isVerticalWin := board[i][0] == player && areEquals(board[i][0], board[i][1], board[i][2])
		if isHorizontalWin || isVerticalWin {
			return true
		}
	}

	isDiagonal1Win := areEquals(board[0][0], board[1][1], board[2][2])
	isDiagonal2Win := areEquals(board[0][2], board[1][1], board[2][0])
	return board[1][1] == player && (isDiagonal1Win || isDiagonal2Win)
}

func isGameDraw() bool {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == "" {
				return false
			}
		}
	}

	return true
}
