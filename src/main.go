package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"

	"./queens"
	"./util"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if (len(os.Args) > 1 && os.Args[1] == "benchmark") {
		bench := benchmark(70, 5)
		util.WriteResult("donnees.txt",bench)
		fmt.Println("Benchmark terminé !")
		return
	}

	if len(os.Args) < 7 {
		fmt.Println("./genetic_queens <N> <Nbr individus> <Taux mutation> <Taux selection meilleurs> <taux selection pires> <taux selection enfants>")
		return
	}

	N, err := strconv.Atoi(os.Args[1])
	populationCount, err := strconv.Atoi(os.Args[2])
	mutationRate, err := strconv.ParseFloat(os.Args[3], 64)
	bestKeepRate, err := strconv.ParseFloat(os.Args[4], 64)
	worstKeepRate, err := strconv.ParseFloat(os.Args[5], 64)
	childKeepRate, err := strconv.ParseFloat(os.Args[6], 64)

	if err != nil {
		fmt.Println(err)
		return
	}
	
	method := queens.QueensMethod{
		N:               N,
		PopulationCount: populationCount,
		MutationRate:    mutationRate,
		BestKeepRate:    bestKeepRate,
		WorstKeepRate:   worstKeepRate,
		ChildKeepRate:   childKeepRate}
	// mesure du temps d'exécution
	start := time.Now()

	result, generation := startGeneticAlgorithm(method)

	duration := time.Since(start)
	fmt.Println("Résultat trouvé en ", duration)
	fmt.Println("Au bout de", generation, "générations")
	fmt.Println(result.Board)
}

func benchmark(target int, interval int) string {
	n := interval
	str := ""

	for (n <= target) {
		method := queens.QueensMethod{
			N:               n,
			PopulationCount: 10000,
			MutationRate:    0.05,
			BestKeepRate:    0.15,
			WorstKeepRate:   0.05,
			ChildKeepRate:   0.8}
		// mesure du temps d'exécution
		start := time.Now()
	
		startGeneticAlgorithm(method)
	
		duration := time.Since(start)
		str += fmt.Sprint(n, duration.Seconds(), "\n")

		n += interval
	}

	return str
}

func startGeneticAlgorithm(method queens.QueensMethod) (queens.QueenSolution, int) {
	population := method.Populate()
	currentGeneration := 0

	var result queens.QueenSolution
	var children []queens.QueenSolution
	var newPop []queens.QueenSolution

	// tri de la population initiale par score
	sort.Slice(population, func(i, j int) bool {
		return population[i].Score > population[j].Score
	})
	result = population[0] // result = meilleur individu

	for result.Score < method.N {
		currentGeneration++

		fmt.Printf("(%d / %d) : génération %d \n", result.Score-1, method.N, currentGeneration)
		children = make([]queens.QueenSolution, 0)
		newPop = make([]queens.QueenSolution, 0)

		// reproduction
		pairs := method.Selection(population)
		for _, pair := range pairs {
			child := pair.Crossover()

			if rand.Float64() <= method.MutationRate {
				child.Mutate()
			}

			children = append(children, child)
		}

		// génération de la nouvelle population à partir des enfants, meilleurs et pires individus
		indexChildren := int((1 - method.ChildKeepRate) * float64(len(children)))
		indexBest := int(float64(len(population)) * method.BestKeepRate)
		indexWorst := int(float64(len(population)) * (1 - method.WorstKeepRate))

		// on garde 80% des enfants
		newPop = append(newPop, children[indexChildren:]...)
		// 15% des meilleurs individus de cette génération
		newPop = append(newPop, population[:indexBest]...)
		// 5% des pires
		newPop = append(newPop, population[indexWorst:]...)

		population = newPop

		// tri de la nouvelle population par fitness
		sort.Slice(population, func(i, j int) bool {
			return population[i].Score > population[j].Score
		})

		result = population[0]
	}

	return result, currentGeneration
}

