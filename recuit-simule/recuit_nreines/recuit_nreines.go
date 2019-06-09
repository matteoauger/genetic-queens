package recuit_nreines

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
)

// function type
type fn func([]int) float64

// boundary
const boundary float64 = 1e-41
const descente float64 = .991
const nbIter int = 100

// Recuit simulé pour N reines
func Recuit(X0 []int, F fn, filename string) []int {
	str := ""
	i := 0
	X := X0
	T := decroissance(10)

	for F(X) > boundary {
		for m := 0; m < nbIter; m++ {
			Y := voisin(X)
			dF := F(Y) - F(X)

			if accepte(dF, T) {
				X = Y
			}
		}

		T = decroissance(T)

		fmt.Println(i, T, F(X))
		str += fmt.Sprint(i, T, F(X), "\n")
		i++
	}

	//writeResult(filename, str)
	return X
}

func accepte(dF float64, T float64) bool {
	if dF >= 0 {
		A := math.Exp(-dF / T)

		if rand.Float64() >= A {
			return false
		}
	}

	return true
}

func voisin(X []int) []int {
	var voisin []int
	var case1, case2, tmp int

	// on copie le tableau X dans le tableau local "voisins"
	// afin que cette fonction soit pure
	voisin = make([]int, len(X))
	copy(voisin, X)

	// on choisit deux cases aléatoirement dans le tableau
	case1 = rand.Intn(len(voisin))
	case2 = rand.Intn(len(voisin))

	// on permute le contenu de case1 avec le contenu de case2
	tmp = voisin[case1]
	voisin[case1] = voisin[case2]
	voisin[case2] = tmp

	return voisin
}

func decroissance(temp float64) float64 {
	//res := temp - descente
	res := temp * descente
	if res >= 0 {
		return res
	} else {
		return 0
	}
}

// écrit la string content dans le fichier fileName
func WriteResult(filename string, content string) {
	byteArr := []byte(content)
	err := ioutil.WriteFile(filename, byteArr, 0644)

	// gestion de l'erreur
	if err != nil {
		panic(err)
	}
}
