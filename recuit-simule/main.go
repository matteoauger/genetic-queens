package main

import (
	"fmt"
	"math"
	"math/rand"

	"time"

	"./recuit_nreines"
)

func f(x float64) float64 {
	return 10 * math.Sin((0.3*x)*math.Sin(1.3*math.Pow(x, 2)+0.00001*math.Pow(x, 4)+0.2*x+80))
}

// Retourne le nombre de reines en prise dans le tableau
func qualite(plateau []int) float64 {
	result := 0.0

	for index, element := range plateau {
		if enPrise(plateau, index, element) {
			result = result + 1.0
		}
	}

	return result
}

// retourne si une reine est en prise
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

func initPlateau(n int) []int {
	slice := make([]int, n)

	for i := 0; i < n; i++ {
		slice[i] = i
	}

	return slice
}

func main() {
	//result := recuit.Recuit(0, f)
	//fmt.Printf("X = %f, Y = %f", result, f(result))
	
/*
	if len(os.Args) <= 0 {
		panic("USAGE : go run main.go [nb iterations]")
	}

	rand.Seed(time.Now().UnixNano())

	n, _ := strconv.Atoi(os.Args[1])
	plateau := initPlateau(n)
	resultNreines := recuit_nreines.Recuit(plateau, qualite, "results.txt")

	fmt.Println("\n\nRésultats :")
	fmt.Println(resultNreines)
	fmt.Println(qualite(resultNreines))

	fmt.Println()
*/

	rand.Seed(time.Now().UnixNano())
	n := 5
	str := ""
	
	for n <= 70 {
		plateau := initPlateau(n)

		start := time.Now()
	
		recuit_nreines.Recuit(plateau, qualite, "results.txt")
	
		duration := time.Since(start)
		str += fmt.Sprint(n, duration.Seconds(), "\n")

		n += 5
	}

	recuit_nreines.WriteResult("donnees.txt", str)
}
