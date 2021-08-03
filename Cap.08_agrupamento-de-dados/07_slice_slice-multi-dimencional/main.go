package main

import "fmt"

/*
- Slices multi-dimensionais são slices que contem slices.
- São como planilhas.
- [][]type
*/

func main() {
	ss := [][]int{
		// Índice:0  1  2                  // Índice:
		[]int{1, 2, 3, 4, 5, 6},       // 0
		[]int{7, 8, 9, 10, 11, 12},    // 1
		[]int{13, 14, 15, 16, 17, 18}, // 2
	}
	fmt.Println(ss[2][4])
}
