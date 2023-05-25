package main

import "fmt"

func main() {
	var idade int64
	fmt.Println("insira a idade: ")
	fmt.Scan(&idade)
	if idade <= 9 {
		fmt.Println("mirim")
	}
	if idade >= 10 && idade <= 13 {
		fmt.Println("infantil")
	}
	if idade >= 14 && idade <= 17 {
		fmt.Println("juvenil")
	}
	if idade > 18 {
		fmt.Println("adulto")
	}
}
