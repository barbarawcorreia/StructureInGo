package main

import "fmt"

func main() {

}

// LIFO - LAST IN FIRST OUT
func vagoesLifo(valoresEntrada []int) []int {
	// resposta := ""
	valoresSaida := []int{}

	for index := len(valoresEntrada) - 1; index >= 0; index-- {
		fmt.Println(index)
		valoresSaida = append(valoresSaida, valoresEntrada[index])
		fmt.Println("valoresSaida", valoresSaida)
	}

	return valoresSaida
}
