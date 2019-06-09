package recuit

import(
	"math/rand"
	"math"
)

// function type
type fn func(float64) float64

// consts
const ESPILON  float64 = 1e-8
const BOUNDARY float64 = -10 + ESPILON
const DESCENTE float64 = 0.87
const NB_ITER  int     = 50

func Recuit(X0 float64, F fn) float64 {
	X  := X0
	T  := Decroissance(100)
	Nt := NB_ITER

	for F(X) >= BOUNDARY {
		for m := 0; m < Nt; m++ {
			Y  := Voisin(X)
			dF := F(Y) - F(X)
			
			if Accepte(dF, T) {
				X = Y
			}
		}

		T = Decroissance(T)
	}
	
	return X
}

func Accepte(dF float64, T float64) bool {
	if dF >= 0 {
		A := math.Exp( -dF / T )
		
		if rand.Float64() >= A {
			return false
		}
	}

	return true;
}

func Voisin(X float64) float64 {
	return X + rand.NormFloat64()
}

func Decroissance(temp float64) float64 {
	return temp - DESCENTE;
}