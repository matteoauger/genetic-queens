package queens

import "math/rand"

type QueensMethod struct {
	N               int
	PopulationCount int
	MutationRate    float64
	BestKeepRate    float64
	WorstKeepRate   float64
	ChildKeepRate   float64
}

func (q QueensMethod) Populate() []QueenSolution {
	res := make([]QueenSolution, q.PopulationCount)

	for i := 0; i < q.PopulationCount; i++ {
		res[i] = CreateRandomSolution(q.N, 100)
	}

	return res
}

func (q QueensMethod) Selection(population []QueenSolution) []QueenSolutionPair {
	var pairs []QueenSolutionPair
	nbPairs := len(population)
	sum := 0
	// calcul du score normalisé
	for _, elt := range population {
		sum += elt.Score
	}
	for i := 0; i < len(population); i++ {
		population[i].NormalizedScore = (float64)(population[i].Score) / (float64)(sum)
	}

	for i := 0; i < nbPairs; i++ {
		parent1 := population[roulette(population)]
		parent2 := population[roulette(population)]

		pairs = append(pairs, QueenSolutionPair{A: parent1, B: parent2})
	}

	return pairs
}

func roulette(population []QueenSolution) int {
	var selected int

	// pick un individu
	rand := rand.Float64()
	normalizedSum := 0.0

	for i := 0; i < len(population); i++ {
		normalizedSum += population[i].NormalizedScore
		if normalizedSum > rand {
			selected = i
			break
		}
	}

	// suppression de l'individu
	return selected
}
