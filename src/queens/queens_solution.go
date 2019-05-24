package queens

import (
	"math"
	"math/rand"
)

type QueenSolution struct {
	Board           []int
	NormalizedScore float64
	Score           int
}

func (solution QueenSolution) Fitness() int {
	n := len(solution.Board)
	result := 0

	for index, element := range solution.Board {
		if enPrise(solution.Board, index, element) {
			result = result + 1
		}
	}

	return n - result + 1
}

func (solution *QueenSolution) Mutate() {
	board := make([]int, len(solution.Board))
	n := len(board)
	copy(board, solution.Board)

	i := rand.Intn(n)
	j := rand.Intn(n)

	// on fait attention à ce que la mutation soit réellement effective (que les 2 gènes soient différents !)
	for j == i {
		j = rand.Intn(n)
	}

	// permutation
	board[i], board[j] = board[j], board[i]

	// mise à jour du plateau
	solution.Board = board
}

func enPrise(plateau []int, i int, j int) bool {
	for k := 0; k < len(plateau); k++ {
		if plateau[k] != j {
			ei := (int)(math.Abs((float64)(i - k)))
			ej := (int)(math.Abs((float64)(j - plateau[k])))

			if ei == ej {
				return true
			}
		}
	}

	return false
}
