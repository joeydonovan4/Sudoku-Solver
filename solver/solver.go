package solver

// Backtracking algorithm for sudoku-solving
func (b *Board) Backtrack() bool {
	// Finds first empty cell on board
	row, col, isEmpty := b.findUnsolvedCell()
	if !isEmpty {
		return true // If no empty cells, then the board must be complete
	}
	for i := 1; i < 10; i++ { // For all possible values
		// Check if possible values are acceptable in the empty cell
		if b.isValidNumber(row, col, i) {
			b.Cells[row][col].solved = true
			b.Cells[row][col].Value = i

			// Recursively call backtrack algorithm
			if b.Backtrack() {
				return true
			}

			// Reset cell to unsolved
			b.Cells[row][col].solved = false
			b.Cells[row][col].Value = 0
		}
	}
	return false
}

// Searches board for an empty Cell
func (b *Board) findUnsolvedCell() (int, int, bool) {
	for _, r := range b.Cells {
		for _, c := range r {
			if c.solved == false {
				return c.Row, c.Col, true
			}
		}
	}
	return 0, 0, false
}

// Checks if the given num is valid
// Valid if num is not within same row, col, and 3x3 box
func (b *Board) isValidNumber(row, col, num int) bool {
	if !b.isValidRow(row, num) ||
		!b.isValidColumn(col, num) ||
		!b.isValidBox(row-row%3, col-col%3, num) {
		return false
	}
	return true
}

// Checks if the given num is valid in the given row on the board
// Returns true if valid, false otherwise
func (b *Board) isValidRow(row, num int) bool {
	for _, c := range b.Cells[row] {
		if c.Value == num {
			return false
		}
	}
	return true
}

// Checks if the given num is valid in the given col on the board
// Returns true if valid, false otherwise
func (b *Board) isValidColumn(col, num int) bool {
	for i := 0; i < len(b.Cells); i++ {
		if b.Cells[i][col].Value == num {
			return false
		}
	}
	return true
}

// Checks if the given num is valid in the corresponding 3x3 box
// Returns true if valid, false otherwise
func (b *Board) isValidBox(startRow, startCol, num int) bool {
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if b.Cells[r+startRow][c+startCol].Value == num {
				return false
			}
		}
	}
	return true
}
