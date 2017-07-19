package solver

import (
	"testing"
)

var (
	cell1  Cell  = Cell{Col: 0, Row: 0, solved: true, Value: 1}
	cell2  Cell  = Cell{Col: 0, Row: 1, solved: true, Value: 2}
	cell3  Cell  = Cell{Col: 0, Row: 2, solved: true, Value: 3}
	cell4  Cell  = Cell{Col: 1, Row: 0, solved: true, Value: 4}
	cell5  Cell  = Cell{Col: 1, Row: 1, solved: true, Value: 5}
	cell6  Cell  = Cell{Col: 1, Row: 2, solved: true, Value: 6}
	cell7  Cell  = Cell{Col: 2, Row: 0, solved: true, Value: 7}
	cell8  Cell  = Cell{Col: 2, Row: 1, solved: true, Value: 8}
	cell9  Cell  = Cell{Col: 2, Row: 2, solved: true, Value: 9}
	cell10 Cell  = Cell{2, 2, false, 0}
	board1 Board = Board{[][]Cell{
		[]Cell{cell1, cell4, cell7},
		[]Cell{cell2, cell5, cell8},
		[]Cell{cell3, cell6, cell10}}}
	board2 Board = Board{[][]Cell{
		[]Cell{cell1, cell4, cell7},
		[]Cell{cell2, cell5, cell8},
		[]Cell{cell3, cell6, cell9}}}
	board3 Board = Board{[][]Cell{
		[]Cell{Cell{0, 0, true, 5}, Cell{1, 0, true, 1}, Cell{2, 0, true, 6},
			Cell{3, 0, true, 8}, Cell{4, 0, true, 4}, Cell{5, 0, true, 9},
			Cell{6, 0, true, 7}, Cell{7, 0, true, 3}, Cell{8, 0, true, 2}},
		[]Cell{Cell{0, 1, true, 3}, Cell{1, 1, false, 0}, Cell{2, 1, true, 7},
			Cell{3, 1, true, 6}, Cell{4, 1, false, 0}, Cell{5, 1, true, 5},
			Cell{6, 1, false, 0}, Cell{7, 1, false, 0}, Cell{8, 1, false, 0}},
		[]Cell{Cell{0, 2, true, 8}, Cell{1, 2, false, 0}, Cell{2, 2, true, 9},
			Cell{3, 2, true, 7}, Cell{4, 2, false, 0}, Cell{5, 2, false, 0},
			Cell{6, 2, false, 0}, Cell{7, 2, true, 6}, Cell{8, 2, true, 5}},
		[]Cell{Cell{0, 3, true, 1}, Cell{1, 3, true, 3}, Cell{2, 3, true, 5},
			Cell{3, 3, false, 0}, Cell{4, 3, true, 6}, Cell{5, 3, false, 0},
			Cell{6, 3, true, 9}, Cell{7, 3, false, 0}, Cell{8, 3, true, 7}},
		[]Cell{Cell{0, 4, true, 4}, Cell{1, 4, true, 7}, Cell{2, 4, true, 2},
			Cell{3, 4, true, 5}, Cell{4, 4, true, 9}, Cell{5, 4, true, 1},
			Cell{6, 4, false, 0}, Cell{7, 4, false, 0}, Cell{8, 4, true, 6}},
		[]Cell{Cell{0, 5, true, 9}, Cell{1, 3, true, 6}, Cell{2, 3, true, 8},
			Cell{3, 5, true, 3}, Cell{4, 5, true, 7}, Cell{5, 5, false, 0},
			Cell{6, 5, false, 0}, Cell{7, 5, true, 5}, Cell{8, 5, false, 0}},
		[]Cell{Cell{0, 6, true, 2}, Cell{1, 6, true, 5}, Cell{2, 6, true, 3},
			Cell{3, 6, true, 1}, Cell{4, 6, true, 8}, Cell{5, 6, true, 6},
			Cell{6, 6, false, 0}, Cell{7, 6, true, 7}, Cell{8, 6, true, 4}},
		[]Cell{Cell{0, 7, true, 6}, Cell{1, 7, true, 8}, Cell{2, 7, true, 4},
			Cell{3, 7, true, 2}, Cell{4, 7, false, 0}, Cell{5, 7, true, 7},
			Cell{6, 7, true, 5}, Cell{7, 7, false, 0}, Cell{8, 7, false, 0}},
		[]Cell{Cell{0, 8, true, 7}, Cell{1, 8, true, 9}, Cell{2, 8, true, 1},
			Cell{3, 8, false, 0}, Cell{4, 8, true, 5}, Cell{5, 8, false, 0},
			Cell{6, 8, true, 6}, Cell{7, 8, false, 0}, Cell{8, 8, true, 8}},
	}}
)

var backtrackResults = []struct {
	board    Board
	expected bool
}{
	{board1, true},
	{board2, true},
	{board3, false},
}

var unsolvedCellResults = []struct {
	board    Board
	row, col int
	b        bool
}{
	{board1, 2, 2, true},
	{board2, 0, 0, false},
}

var validNumberResults = []struct {
	board         Board
	row, col, num int
	expected      bool
}{
	{board1, 2, 2, 9, true},
	{board1, 2, 2, 8, false},
}

var validRowResults = []struct {
	board    Board
	row, num int
	expected bool
}{
	{board1, 1, 3, true},
	{board1, 0, 1, false},
}

var validColumnResults = []struct {
	board    Board
	col, num int
	expected bool
}{
	{board1, 0, 4, true},
	{board1, 1, 5, false},
}

var validBoxResults = []struct {
	board         Board
	row, col, num int
	expected      bool
}{
	{board1, 0, 0, 7, false},
	{board1, 0, 0, 9, true},
}

func TestBackTrack(t *testing.T) {
	for _, results := range backtrackResults {
		actual := results.board.Backtrack()
		if actual != results.expected {
			t.Errorf("Backtrack(): expected %t, actual %t", results.expected, actual)
		}
	}
}

func TestFindUnsolvedCell(t *testing.T) {
	for _, results := range unsolvedCellResults {
		row, col, b := results.board.findUnsolvedCell()
		if row != results.row && col != results.col && b != results.b {
			t.Errorf("format")
		}
	}
}

func TestIsValidNumber(t *testing.T) {
	for _, results := range validNumberResults {
		actual := results.board.isValidNumber(results.row, results.col, results.num)
		if actual != results.expected {
			t.Errorf("format")
		}
	}
}

func TestIsValidRow(t *testing.T) {
	for _, results := range validRowResults {
		actual := results.board.isValidRow(results.row, results.num)
		if actual != results.expected {
			t.Errorf("format")
		}
	}
}

func TestIsValidColumn(t *testing.T) {
	for _, results := range validColumnResults {
		actual := results.board.isValidColumn(results.col, results.num)
		if actual != results.expected {
			t.Errorf("format")
		}
	}
}

func TestIsValidBox(t *testing.T) {
	for _, results := range validBoxResults {
		actual := results.board.isValidBox(results.row, results.col, results.num)
		if actual != results.expected {
			t.Errorf("format")
		}
	}
}
