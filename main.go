package main

import (
	"fmt"
)

const alive string = "X"
const dead string = "0"

func board() [][]string {
	row1 := []string{dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead}
	row2 := []string{dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead}
	row3 := []string{dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead}
	row4 := []string{dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead}
	row5 := []string{dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead}
	row6 := []string{dead, dead, dead, dead, dead, dead, dead, dead, alive, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead}
	row7 := []string{dead, dead, dead, dead, dead, dead, dead, dead, alive, alive, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead}
	row8 := []string{dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead}
	row9 := []string{dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead}
	row10 := []string{dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead}

	return [][]string{row1, row2, row3, row4, row5, row6, row7, row8, row9, row10}
}

func printBoard(board [][]string) {
	for i := 0; i < len(board); i++ {
		fmt.Println(board[i])
	}

}

func main() {
	printBoard(board())
}
