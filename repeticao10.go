package main

import (
	"fmt"
	"sort"
)

func main() {
	var numeros int
	lista := []int{}
	for {
		fmt.Println("Insira um número, para parar digite 0 ")
		fmt.Scan(&numeros)
		if numeros == 0 {
			break
		}
		lista = append(lista, numeros)
	}
	sort.Ints(lista)
	fmt.Print("O maior núemro é: ", lista[len(lista)-1])
}
