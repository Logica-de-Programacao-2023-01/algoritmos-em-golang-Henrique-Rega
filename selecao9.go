package main

import (
	"fmt"
	"sort"
)

func main() {
	var x, y, z float64
	fmt.Scanln(&x)
	fmt.Scanln(&y)
	fmt.Scanln(&z)
	numeros := []float64{x, y, z}
	sort.Float64s(numeros)
	fmt.Println(numeros[0], numeros[1], numeros[2])
}
