package main

import "fmt"

func main() {
	for i := 0; i < 22; i++ {
		for j := i; j%2 <= 0; j++ {
			fmt.Println(j)
		}
	}
}
