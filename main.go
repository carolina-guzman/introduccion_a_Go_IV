package main

import "fmt"

func fibonacci(i int, x int, res []int) int {
	if i > x {
		fmt.Println(res)
		return 0
	}
	if i == 0 { // caso base
		res = append(res, 0)
		res = append(res, 1)
		i++
	} else {
		res = append(res, res[i-1]+res[i-2])
	}

	return fibonacci(i+1, x, res)
}

func main() {
	var res []int
	fibonacci(0, 10, res)
}
