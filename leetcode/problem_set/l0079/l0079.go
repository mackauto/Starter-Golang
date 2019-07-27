package main

import "fmt"

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if DFS(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func DFS(board [][]byte, i, j int, word string, index int) bool {
	if index == len(word) {
		return true
	}
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || word[index] != board[i][j] {
		return false
	}
	temp := board[i][j]
	board[i][j] = '%'

	r := DFS(board, i+1, j, word, index+1) || DFS(board, i, j+1, word, index+1) || DFS(board, i-1, j, word, index+1) || DFS(board, i, j-1, word, index+1)
	board[i][j] = temp
	return r
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, "ABCCED"))
	fmt.Println(exist(board, "SEE"))
	fmt.Println(exist(board, "ABCB"))
}
