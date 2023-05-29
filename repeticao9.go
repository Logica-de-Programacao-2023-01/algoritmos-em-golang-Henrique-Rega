package main

import "fmt"

func main() {
	var valores, soma, numero int64
	for {
		fmt.Println("Digite os números para o cálculo da média(digite 0 quando acabar de inserir os números): ")
		fmt.Scan(&valores)
		if valores == 0 {
			break
		}
		soma += valores
		numero++
	}
	if numero > 0 {
		media := soma / numero
		fmt.Println("a méida é: ", media)
	}
}
