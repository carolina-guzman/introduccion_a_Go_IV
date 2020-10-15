package main

import "fmt"

func fibonacci(i int, x int, res []int) int {
	if i == x {
		fmt.Println("Resultado: ", res)
		return 0
	}
	if i == 0 { // caso base
		res = append(res, 0, 1)
		i++
	} else {
		res = append(res, res[i-1]+res[i-2])
	}

	return fibonacci(i+1, x, res)
}

func NumMayor(args ...int) int {
	mayor := 0
	for _, v := range args {
		if v > mayor {
			mayor = v
		}
	}

	return mayor
}

func generadorImpares() func() uint {
	i := uint(1)
	return func() uint {
		var par = i
		i += 2
		return par
	}
}

func intercambia(a *int, b *int)  { 
	*a , *b = *b , *a 
}

func main() {
	var res []int
	var n int
	fmt.Println("\n--------------FIBONACCI----------------")
	fmt.Println("Ingrese el la cantidad de numeros para la serie Fibonacci:")
	fmt.Scan(&n)
	fibonacci(0, n, res)

	fmt.Println("\n--------------Numero mayor----------------")
	a := []int{1,5,8,9,2,6,19,0}
	fmt.Println("El numero mayor de: ", a, "es: ", NumMayor(a...))

	fmt.Println("\n--------------Numeros Impares----------------")
	nextPar := generadorImpares()
	fmt.Println(nextPar())
	fmt.Println(nextPar())
	fmt.Println(nextPar())
	fmt.Println(nextPar())
	fmt.Println(nextPar())
	fmt.Println(nextPar())
	fmt.Println(nextPar())

	fmt.Println("\n--------------Interbambiar numeros----------------")
	a1 := 10
	b1 := 3
	fmt.Println("a: ", a1, ", b: ", b1)
	intercambia(&a1, &b1)
	fmt.Println("a: ", a1, ", b: ", b1)

	
}
