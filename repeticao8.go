package main

import "fmt"

func main() {
	x := 0
	fmt.Println("escreva o número")
	fmt.Scanln(&x)
	for i := 1; i < x+1; i++ {
		for j := i; x%j == 0; j++ {
			fmt.Println(j)
		}

	}
}
