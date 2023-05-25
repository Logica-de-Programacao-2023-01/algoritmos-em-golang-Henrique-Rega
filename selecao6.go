package main

import "fmt"

func main() {
	var x int64
	var y int64
	fmt.Println("insira os números a serem multiplicados ou somados")
	fmt.Scanln(&x, &y)
	if y == 0 {
		fmt.Println(x)
	} else if y > 0 && x > 0 || y < 0 && x < 0 {
		fmt.Println(x * y)
	} else {
		fmt.Println(x + y)
	}
}
