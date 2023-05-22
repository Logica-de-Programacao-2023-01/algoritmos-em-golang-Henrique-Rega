package main

import (
	"fmt"
	"sort"
)

func main() {
	x := 0
	y := 0
	z := 0
	fmt.Scan(&x, &y, &z)
	a := []int{}
	a = append(a, x, y, z)
	sort.Ints(a)
	fmt.Println(a[0])
}
