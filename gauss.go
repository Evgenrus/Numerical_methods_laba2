package main

import (
	"math"
)

func gauss(a [][]float64, y []float64, i int, j int, eps float64) []float64 {
	result := make([]float64, i)

	tef := make([][]float64, i)
	for l := 0; l < i; l++ {
		tef[l] = make([]float64, i)
	}

	copy(tef, a)
	//det := determinant(tef, i)
	//println("определитель матрицы: ", det)
	//println("Обратная матрица: ", reverseMatrix(tef, i, det))

	var k, index int = 0, 0
	var max float64 = 0
	for ; k < i; {
		max = a[k][k]
		index = k
		for l := k+1; l < i; l++ {
			if math.Abs(a[l][k]) > math.Abs(max) {
				max = a[l][k]
				index = l
			}
		}
		//changing position of equations
		for l := 0; l < i; l++ {
			var temp float64 = a[k][l]
			a[k][l] = a[index][l]
			a[index][l] = temp
		}
		var temp float64 = y[k]
		y[k] = y[index]
		y[index] = temp
		//normalize
		for l := k; l < i; l++ {
			temp = a[l][k]
			if math.Abs(temp) < eps {
				continue
			}
			for u := 0; u < j; u++ {
				a[l][u] /= temp
			}
			y[l] /= temp
			//subtract equations
			if l == k {
				continue
			}
			for u := 0; u < j; u++ {
				a[l][u] -= a[k][u]
			}
			y[l] -= y[k]
		}
		k++
	}
	//calculating variables
	for l := i-1; l >= 0; l-- {
		result[l] = y[l]
		for u := 0; u < l; u++ {
			y[u] -= a[u][l] * result[l]
		}
	}

	return result
}
