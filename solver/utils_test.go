package solver

import (
	"fmt"
	"testing"
)

var fillBoardResults = []struct {
	path     string
	expected string
}{
	{"exampleBoard3.txt", board3.String()},
}

var generateCellResults = []struct {
	col, row int
	val      string
	expected Cell
}{
	{0, 0, "_", Cell{0, 0, false, 0}},
	{1, 7, "2", Cell{1, 7, true, 2}},
	{3, 3, "$", Cell{3, 3, false, 0}},
}

var generateRowResults = []struct {
	fields   []string
	row      int
	expected []Cell
}{
	{[]string{"1", "_", "3", "7"}, 3,
		[]Cell{
			Cell{0, 3, true, 1}, Cell{1, 3, false, 0}, Cell{2, 3, true, 3}, Cell{3, 3, true, 7},
		},
	},
	{[]string{"$"}, 1, []Cell{}},
}

var cellStringResults = []struct {
	cell     Cell
	expected string
}{
	{Cell{1, 7, false, 0}, "_"},
	{Cell{5, 1, true, 1}, "1"},
	{Cell{8, 3, true, 5}, "5"},
}

//const string result =
var boardStringResults = []struct {
	board    Board
	expected string
}{
	{board1, "147\n258\n36_\n\n"},
	{board3, fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n",
		"516|849|732", "3_7|6_5|___", "8_9|7__|_65\n",
		"135|_6_|9_7", "472|591|__6", "968|37_|_5_\n",
		"253|186|_74", "684|2_7|5__", "791|_5_|6_8\n")},
}

func TestFillBoard(t *testing.T) {
	for _, results := range fillBoardResults {
		actual, _ := FillBoard(results.path)
		fmt.Printf(actual.String())
		if actual.String() != results.expected {
			t.Errorf("FillBoard(): expected %s, actual %s",
				results.expected, actual.String())
		}
	}
}

func TestGenerateRow(t *testing.T) {
	for _, results := range generateRowResults {
		actual, _ := generateRow(results.fields, results.row)
		for i := 0; i < len(actual); i++ {
			if actual[i] != results.expected[i] {
				t.Errorf("generateRow(%v, %d): expected %v, actual: %v", results.fields,
					results.row, results.expected, actual)
			}
		}
	}
}

func TestGenerateCell(t *testing.T) {
	for _, results := range generateCellResults {
		actual, _ := generateCell(results.col, results.row, results.val)
		if actual != results.expected {
			t.Errorf("generateCell(%d, %d, %s): expected %d, actual %d",
				results.col, results.row, results.val, results.expected, actual)
		}
	}
}

func TestCellString(t *testing.T) {
	for _, results := range cellStringResults {
		actual := results.cell.String()
		if actual != results.expected {
			t.Errorf("String(): expected %s, actual %s", results.expected, actual)
		}
	}
}

func TestBoardString(t *testing.T) {
	for _, results := range boardStringResults {
		actual := results.board.String()
		if actual != results.expected {
			t.Errorf("String(): expected %s, actual %s", results.expected, actual)
		}
	}
}
