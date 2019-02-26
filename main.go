package main

import "fmt"

func main() {

}

// LIFO - LAST IN FIRST OUT
func vagoesLifo(valoresEntrada []int) []int {
	valoresSaida := []int{}

	for index := len(valoresEntrada) - 1; index >= 0; index-- {
		fmt.Println(index)
		valoresSaida = append(valoresSaida, valoresEntrada[index])
		fmt.Println("valoresSaida", valoresSaida)
	}

	return valoresSaida
}

// FIFO - FIRST IN FIRST OUT
func vagoesFifo(valoresEntrada []int) []int {
	valoresSaida := []int{}

	for index := 0; index <= len(valoresEntrada)-1; index++ {
		fmt.Println(index)
		valoresSaida = append(valoresSaida, valoresEntrada[index])
		fmt.Println("valoresSaida", valoresSaida)
	}

	return valoresSaida
}
