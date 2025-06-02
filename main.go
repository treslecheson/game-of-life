package main

import (
	"fmt"
	"strconv"
)

const alive string = "X"
const dead string = "0"

type Board [][]string

func contains(row []string, target string) int {
	count := 0
	for i := 0; i < len(row); i++ {
		if row[i] == target {
			count++
		}
	}
	return count
}

func board() Board {
	row1 := []string{dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead}
	row2 := []string{dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead}
	row3 := []string{dead, dead, alive, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead}
	row4 := []string{dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead}
	row5 := []string{dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead}
	row6 := []string{dead, dead, dead, dead, dead, dead, dead, dead, alive, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead}
	row7 := []string{dead, dead, dead, dead, dead, dead, dead, dead, alive, alive, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead}
	row8 := []string{dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead}
	row9 := []string{dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead}
	row10 := []string{dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead}

	return [][]string{row1, row2, row3, row4, row5, row6, row7, row8, row9, row10}
}

func solitude(board Board) Board {
	for i := 0; i < len(board); i++ {
		if contains(board[i], alive) == 1 {
			if contains(board[i-1], alive) == 0 && contains(board[i+1], alive) == 0 {
				fmt.Println("Row " + strconv.Itoa(i) + " died")
			}

		}

	}

	return board

}

func printBoard(board [][]string) {
	for i := 0; i < len(board); i++ {
		fmt.Println(board[i])
	}

}

func main() {
	printBoard(board())
	solitude(board())
}
