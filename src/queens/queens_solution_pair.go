package queens

import (
	"math"
	"math/rand"
	"sort"
)

type QueenSolutionPair struct {
	A QueenSolution
	B QueenSolution
}

func (pair QueenSolutionPair) Crossover() QueenSolution {
	n := len(pair.A.Board)
	blines := make([]int, len(pair.B.Board))

	// copie de B.Board
	childBoard := make([]int, len(pair.B.Board))
	copy(childBoard, pair.B.Board)

	// représentation inverse à la représentation actuelle : le tableau stocke les lignes(valeur) dans les colonnes(index)
	for i := 0; i < n; i++ {
		blines[pair.B.Board[i]] = i
	}

	k1 := rand.Intn(n)
	k2 := rand.Intn(n)
	for k1 == k2 { // tq k1 = k2, on tire à nouveau k2
		k2 = rand.Intn(n)
	}

	var b int = int(math.Min(float64(k1), float64(k2)))
	var e int = int(math.Max(float64(k1), float64(k2)))

	lines := make([]int, e-b)
	for i := 0; i < len(lines); i++ {
		lines[i] = blines[pair.A.Board[b+i]]
	}

	sort.Ints(lines)

	for i := 0; i < len(lines); i++ {
		childBoard[lines[i]] = pair.A.Board[b+i]
	}

	return CreateSolution(childBoard)
}
