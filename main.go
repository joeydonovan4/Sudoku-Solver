package main

import (
	"fmt"
	"os"

	"github.com/joeydonovan4/sudoku/solver"
)

// Read board from file
func main() {
	file := os.Args
	if len(file) <= 1 {
		fmt.Printf("Please provide a filename of a sudoku board as a command-line argument\n")
	} else {
		board, err := solver.FillBoard(fmt.Sprintf("%v", file[1]))
		if err != nil {
			fmt.Printf("Error:\n%s", err)
		}
		fmt.Printf("\nOriginal board:\n%v", board)

		if board.Backtrack() {
			fmt.Printf("Solved board:\n%v", board)
		} else {
			// Do something to show that there is no solution for this board
			fmt.Println("No solution for the given board.")
		}
	}
}
