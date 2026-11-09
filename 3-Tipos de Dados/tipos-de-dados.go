package main

import (
	"errors"
	"fmt"
)

func main() {
	// NÚMEROS INTEIROS

	// Existem 4 tipos de números inteiros no Go, que são: int8, int16, int32 e int64.
	// int8, int16, int32, int64 := 127, 32767, 2147483647, 9223372036854775807

	var numero int8 = 100
	fmt.Println(numero)

	var numero1 int16 = 10000
	fmt.Println(numero1)

	var numero2 int32 = 100000
	fmt.Println(numero2)

	var numero3 int64 = 10000000000
	fmt.Println(numero3)

	// uint // unsygned int - Existem 4 tipos de números inteiros sem sinal no Go, que são: uint8, uint16, uint32 e uint64.

	var numero20 uint32 = 10000
	fmt.Println(numero20)

	var numeroSemSinal uint64 = 1000000000000000
	fmt.Println(numeroSemSinal)

	// alias - é usado quando está mechendo com números que representam caracteres, como por exemplo,
	// o tipo byte é um alias para uint8 e o tipo rune é um alias para int32.]
	// INT32 = RUNE
	var numero30 rune = 12456
	fmt.Println(numero30)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	// FIM DOS NÚMEROS INTEIROS

	// NÚMEROS REAIS  - Existem 2 tipos de números reais no Go, que são: float32 e float64. (Números com ponto flutuante "Números quebrados")

	var numeroReal5 float32 = 123.456
	fmt.Println(numeroReal5)

	var numeroReal6 float64 = 123456.7890123456789
	fmt.Println(numeroReal6)

	// FIM DOS NÚMEROS REAIS

	// STRINGS - Uma string é uma sequência de caracteres. No Go, as strings são imutáveis, ou seja, não podem ser alteradas depois de criadas.

	var str string = "Texto1"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	// FIM DAS STRINGS

	texto := 5
	fmt.Println(texto)

	// BOOLEANOS - O tipo booleano é usado para representar valores verdadeiros ou falsos. No Go, o tipo booleano é representado pelo tipo bool.

	var booleano1 bool = true
	fmt.Println(booleano1)

	var booleano2 bool = false
	fmt.Println(booleano2)

	// FIM DOS BOOLEANOS

	// ERROS - O tipo error é usado para representar erros. No Go, o tipo error é representado pelo tipo error.

	var erro error = errors.New("Erro interno") // O "Tipo Error" é muito usado para representar erros. No Go, o tipo error é representado pelo tipo error.
	fmt.Println(erro)                           // Erro interno

}
