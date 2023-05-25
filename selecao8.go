package main

import "fmt"

func main() {
	x := 0
	fmt.Println("escreva o número do dia da semana: ")
	fmt.Scan(&x)
	if x == 1 {
		fmt.Print("domingo")
	} else if x == 2 {
		fmt.Print("segunda-feira")

	} else if x == 3 {
		fmt.Print("terça-feira")

	} else if x == 4 {
		fmt.Print("quarta-feira")

	} else if x == 5 {
		fmt.Print("quinta-feira")

	} else if x == 6 {
		fmt.Print("sexta-feira")

	} else if x == 7 {
		fmt.Print("sábado")

	} else {
		fmt.Println("O número inserido não é correspondente a um dia da semana")
	}
}
