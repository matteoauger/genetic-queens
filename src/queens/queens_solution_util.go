package queens

import (
	"math/rand"
)

func CreateSolution(board []int) QueenSolution {
	solution := QueenSolution{Board: board}
	solution.Score = solution.Fitness()

	return solution
}

func CreateRandomSolution(n int, permutationCount int) QueenSolution {
	board := make([]int, n)

	// construction du plateau
	for i := 0; i < n; i++ {
		board[i] = i
	}

	// on effectue 100 permutations sur le plateau
	for j := 0; j < permutationCount; j++ {
		index1 := rand.Intn(n)
		index2 := rand.Intn(n)

		board[index1], board[index2] = board[index2], board[index1]
	}

	return CreateSolution(board)
}
