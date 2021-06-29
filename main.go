package main

import (
	"fmt"
)

func main() {

	var i, j int
	fmt.Println("Please enter number of equation: ")
	fmt.Scan(&i)
	fmt.Println("Please enter number of variables")
	fmt.Scan(&j)

	var temp, eps float64
	fmt.Println("Please enter eps: ")
	fmt.Scan(&eps)

	//Matrix initialization
	var a [][]float64
	a = make([][]float64, i)
	for k := range a {
		a[k] = make([]float64, j)
	}

	//Matrix fill
	for k := 0; k < i; k++ {
		for l := 0; l < j; l++ {
			fmt.Printf("Please enter [%d][%d]a: ", k, l)
			fmt.Scan(&temp)
			a[k][l] = temp
		}
	}

	var y []float64
	y = make([]float64, i)
	for k := range y {
		fmt.Printf("Please enter [%d]y", k)
		fmt.Scan(&temp)
		y[k] = temp
	}

	fmt.Println(a, y)

	fmt.Println()

	var expr int
	fmt.Println("1. Gauss;\t 2. iterations")
	fmt.Scan(&expr)

	res := make([]float64, i)

	switch expr {
	case 1:
		res = gauss(a, y, i, j, eps)
		fmt.Println(res)
		break
	case 2:
		var er error
		fmt.Println(a, y)
		res, er = iterations(a, y, i, j, eps)
		if er == nil {
			fmt.Println(res)
		} else {
			fmt.Println("error")
		}
		break
	}

}