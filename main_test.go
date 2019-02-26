package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVagoesEntrada1(t *testing.T) {
	// Premissa de que foi dada a quantidade '5' de vagões.
	// qtDeVagoes := 5

	valorEntradaConstante := []int{1, 2, 3, 4, 5}
	valorDaSaídaEsperado := []int{5, 4, 3, 2, 1}
	// Utilizando a estratégia fifo
	// Empilha na estação e depois tira para a direção de B.
	// Quando empilha fica em
	// valorSaidaAposFifo := []int{}
	// respostaEsperada := "yes"

	retornoLIFO := vagoesLifo(valorEntradaConstante)

	assert.Equal(t, valorDaSaídaEsperado, retornoLIFO)

}

// Primeiro vai retornar o valor esperado e depois refatorar pra responder yes or no.
