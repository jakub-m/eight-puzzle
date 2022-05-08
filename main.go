package main

import "fmt"

const (
	boardDim = 3
	gap      = 0
)

type board [boardDim][boardDim]tile
type tile int

var start = board{
	{2, 5, gap},
	{1, 4, 8},
	{7, 3, 6},
}

func (b board) String() string {
	s := ""
	for row := 0; row < boardDim; row++ {
		for col := 0; col < boardDim; col++ {
			t := b[row][col]
			if t == gap {
				s += "."
			} else {
				s += fmt.Sprint(t)
			}
		}
		if row < (boardDim - 1) {
			s += "\n"
		}
	}
	return s
}

func main() {
	// bfs
	visited := make(map[board]bool)
	backtrack := make(map[board]board)
	backlog := []board{start}
	for len(backlog) > 0 {
		curr := backlog[0]
		backlog = backlog[1:]
		if isFinished(curr) {
			// print path
			fmt.Printf("finished\n")
			for curr != start {
				fmt.Printf("%s\n---\n", curr)
				curr = backtrack[curr]
			}
			fmt.Printf("%s\n---\n", curr)
			break
		}
		for _, b := range getNextBoards(curr) {
			if visited[b] {
				continue
			}
			backtrack[b] = curr
			visited[b] = true
			backlog = append(backlog, b)
		}
	}
}

func isFinished(b board) bool {
	return b == board{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, gap},
	}
}

func getNextBoards(curr board) []board {
	next := []board{}
	if b, ok := swap(curr, 1, 0); ok {
		next = append(next, b)
	}
	if b, ok := swap(curr, -1, 0); ok {
		next = append(next, b)
	}
	if b, ok := swap(curr, 0, 1); ok {
		next = append(next, b)
	}
	if b, ok := swap(curr, 0, -1); ok {
		next = append(next, b)
	}
	return next
}

func swap(in board, dr, dc int) (board, bool) {
	ir, ic := findGap(in)
	if (ir+dr) < 0 || (ir+dr) >= boardDim || (ic+dc) < 0 || (ic+dc) >= boardDim {
		return in, false
	}
	in[ir][ic], in[ir+dr][ic+dc] = in[ir+dr][ic+dc], in[ir][ic]
	return in, true
}

func findGap(b board) (int, int) {
	for row := 0; row < boardDim; row++ {
		for col := 0; col < boardDim; col++ {
			if b[row][col] == gap {
				return row, col
			}
		}
	}
	panic("no gap on board")
}
