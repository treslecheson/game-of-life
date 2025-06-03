package main

import (
	"fmt"
)

const alive string = "X"
const dead string = "0"

type Board [][]string

func board() Board {
	row1 := []string{dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead}
	row2 := []string{dead, dead, alive, alive, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead}
	row3 := []string{dead, alive, alive, alive, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead}
	row4 := []string{dead, dead, alive, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead, dead}
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

func contains(row []string, target string) int {
	count := 0
	for i := 0; i < len(row); i++ {
		if row[i] == target {
			count++
		}
	}
	return count
}

func indexes(row []string) []int {
	indicesList := []int{}
	for i := 0; i < len(row); i++ {
		if row[i] == alive {
			indicesList = append(indicesList, i)
		}
	}
	return indicesList
}

func neighbors(checking, cellLow, cellHigh int, above, current, below []string) bool {
	cellsAround := 0
	aboveIdices := indexes(above)
	belowIdices := indexes(below)

	for i := 0; i < len(aboveIdices); i++ {
		if checking == aboveIdices[i] {
			cellsAround++
		} else if checking-1 == aboveIdices[i] {
			cellsAround++
		} else if checking+1 == aboveIdices[i] {
			cellsAround++
		}
	}

	for i := 0; i < len(belowIdices); i++ {
		if checking == belowIdices[i] {
			cellsAround++
		} else if checking-1 == belowIdices[i] {
			cellsAround++
		} else if checking+1 == belowIdices[i] {
			cellsAround++
		}
	}

	if current[checking-1] == alive {
		cellsAround++
	} else if current[checking+1] == alive {
		cellsAround++
	}

	fmt.Println(cellsAround)
	if cellLow <= cellsAround && cellsAround <= cellHigh {
		return true
	}
	return false
}

func solitude(above, below []string) bool {

	if contains(above, alive) == 0 && contains(below, alive) == 0 {
		return true
	}
	return false
}

func populate(checking, cellLow, cellHigh int, above, current, below []string) {
	if neighbors(checking, cellLow, cellHigh, above, current, below) == true {
		current[checking] = alive
	}
}

func kill(checking, cellLow, cellHigh int, above, current, below []string) {
	if neighbors(checking, cellLow, cellHigh, above, current, below) == true {
		current[checking] = dead
	}
}

func main() {
	printBoard(board())
	board := board()

	for i := 0; i < 100; i++ {
		for i := 2; i < len(board)-2; i++ {
			for x := 2; x < len(board[i])-3; x++ {
				if board[i][x] == alive {
					kill(x, 4, 8, board[i-1], board[i], board[i+1])
				} else {
					populate(x, 3, 3, board[i-1], board[i], board[i+1])
					fmt.Println("  ")
					fmt.Println("  ")
					fmt.Println("  ")
					fmt.Println("------------------------------------------------------")
				}
				printBoard(board)
			}

		}

	}

	printBoard(board)

}
