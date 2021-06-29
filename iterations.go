package main

import (
	"errors"
	//"fmt"
	"math"
)

func iterations(a [][]float64, y []float64, i int, j int, eps float64) ([]float64, error) {
	result := make([]float64, i)
	p := make([]float64, i)
	//проверка сходимости
	for l := 0; l < i; l++ {
		if !diagonal(a, l, i) {
			return nil, errors.New("не сходится")
		}
	}
	//начальное приближение
	for l := 0; l < i; l++ {
		p[l] = 0
		result[l] = y[l]
	}
	var iter int = 0
	//итерации
	for ; condition(p, result, i, eps); {
		if iter != 0 {
			for l := 0; l < i; l++ {
				p[l] = result[l]
			}
		}
		//Вычисление x k-тых
		for l := 0; l < i; l++ {
			result[l] = y[l] / a[l][l]
			//fmt.Println(result[l])
			for k := 0; k < j; k++ {
				if k != l {
					result[l] -= a[l][k]*p[k] / a[l][l]
				}
			}
		}
		//fmt.Println(a, y)
		//fmt.Println(result)
		iter++
	}
	return result, nil
}

func diagonal(a[][]float64, i int, size int) bool {
	var sum float64 = 0
	var res bool
	for l := 0; l < size; l++ {
		sum += math.Abs(a[i][l])
	}
	sum -= math.Abs(a[i][i])
	if math.Abs(sum) <= math.Abs(a[i][i]) {
		res = true
	} else {
		res = false
	}
	return res
}

//условие прекращения итераций
func condition(p[]float64, res[]float64, i int, eps float64) bool {
	var max float64 = 0
	for l := 0; l < i; l++ {
		if math.Abs(res[l] - p[l]) > math.Abs(max) {
			max = math.Abs(res[l] - p[l])
		}
	}
	//fmt.Println(max, eps, max <= eps, res[temp], p[temp])
	if max <= eps {
		return false
	} else {
		return true
	}
}