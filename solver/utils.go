package solver

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cell struct {
	Col    int
	Row    int
	solved bool
	Value  int
}

type Board struct {
	Cells [][]Cell
}

// Gets and stores the board data properly
func FillBoard(s string) (Board, error) {
	var b Board
	// Read from file and store cells into proper locations in slice/array
	file, err := os.Open(s)
	if err != nil {
		return b, fmt.Errorf("Error opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		r, err := generateRow(strings.Fields(scanner.Text()), row)
		if err != nil {
			return b, err
		}
		if len(r) > 0 {
			b.Cells = append(b.Cells, r)
			row++
		}
	}

	return b, nil
}

// Creates a corresponding row in the sudoku board
func generateRow(s []string, row int) ([]Cell, error) {
	r := []Cell{}
	for col, val := range s {
		c, err := generateCell(col, row, val)
		if err != nil {
			return r, err
		}
		r = append(r, c)
	}
	return r, nil
}

// Creates a corresponding cell in the sudoku board
func generateCell(col, row int, val string) (Cell, error) {
	c := Cell{Col: col, Row: row}
	if val == "_" {
		c.solved = false
	} else {
		if v, err := strconv.Atoi(val); err == nil {
			c.solved = true
			c.Value = v
		} else {
			return c, err
		}
	}
	return c, nil
}

func (c Cell) String() string {
	if c.Value != 0 {
		return fmt.Sprintf("%v", c.Value)
	} else {
		return "_"
	}
}

func (b Board) String() string {
	res := ""
	for idx, r := range b.Cells {
		for i := 0; i < len(r); i++ {
			res += fmt.Sprintf("%v", r[i])
			if i%3 == 2 && i != len(r)-1 {
				res += "|"
			}
		}
		res += "\n"
		if idx%3 == 2 {
			res += "\n"
		}
	}
	return res
}
