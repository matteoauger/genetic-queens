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

	// on tire deux index au hasard
	rand1 := rand.Intn(n)
	rand2 := rand.Intn(n)
	for rand1 == rand2 { // tq rand1 = rand2, on tire à nouveau rand2 (afin d'avoir deux index bien différents)
		rand2 = rand.Intn(n)
	}

	// on choisit le min et le max afin de distinguer l'index de début de l'index de fin
	var begin int = int(math.Min(float64(rand1), float64(rand2)))
	var end int   = int(math.Max(float64(rand1), float64(rand2)))
	
	lines := make([]int, end-begin)
	for i := 0; i < len(lines); i++ {
		lines[i] = blines[pair.A.Board[begin+i]]
	}
	// tri du tableau
	sort.Ints(lines)

	for i := 0; i < len(lines); i++ {
		childBoard[lines[i]] = pair.A.Board[begin+i]
	}

	return CreateSolution(childBoard)
}
