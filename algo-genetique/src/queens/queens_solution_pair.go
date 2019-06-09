package queens

import (
	//"math"
	"math/rand"
	//"sort"
	"../util"
)

type QueenSolutionPair struct {
	A QueenSolution
	B QueenSolution
}

func (pair QueenSolutionPair) Crossover() QueenSolution {
	n := len(pair.A.Board)
	// on base le plateau de l'enfant sur le plateau du premier parent
	childBoard := make([]int, n)

	copy(childBoard, pair.A.Board)

	randIndex := rand.Intn(len(pair.B.Board))
	childIndex := util.Find(childBoard, pair.B.Board[randIndex])

	childBoard[randIndex], childBoard[childIndex] = childBoard[childIndex], childBoard[randIndex]

	return CreateSolution(childBoard)
}
