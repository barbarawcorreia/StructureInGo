package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

x	func TestVagoesEntrada1(t *testing.T) {
	// Premissa de que foi dada a quantidade '5' de vagões.
	// qtDeVagoes := 5

	valorEntradaConstante := []int{1, 2, 3, 4, 5}
	valorDaSaídaEsperado := []int{5, 4, 3, 2, 1}
	// Utilizando a estratégia Lifo

	// respostaEsperada := "yes"

	retornoLIFO := vagoesLifo(valorEntradaConstante)

	assert.Equal(t, valorDaSaídaEsperado, retornoLIFO)

}

func TestVagoesEntrada2(t *testing.T) {

	valorEntradaConstante := []int{1, 2, 3, 4, 5}
	valorDaSaídaEsperado := []int{1, 2, 3, 4, 5}
	// Utilizando a estratégia fifo - last in first out
	// resposta esperada - yes
	retornoFIFO := vagoesFifo(valorEntradaConstante)

	assert.Equal(t, valorDaSaídaEsperado, retornoFIFO)

}

// Qual estratégia usar (?)
// é obrigado a passar pela estação (?)

// Tentativa de usar FIFO com pulinhos
func TestVagoesEntrada3(t *testing.T) {

	valorEntradaConstante := []int{1, 2, 3, 4, 5, 6}
	valorDaSaídaEsperado := []int{1, 3, 2, 5, 4, 6}
	// Utilizando a estratégia fifo - first in first out
	// resposta esperada - yes
	retornoFIFO := vagoesFifo(valorEntradaConstante)

	assert.Equal(t, valorDaSaídaEsperado, retornoFIFO)

}

//https://www.urionlinejudge.com.br/judge/pt/problems/view/1062
// refatorar p/ atender yes or no da saída do problema.
