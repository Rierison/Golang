package main

import "fmt"

// Função é como se fosse um bloco de código que pode ser chamado em qualquer lugar do programa.
// Ela é usada para organizar o código e torná-lo mais legível e reutilizável.
// É como se fosse uma receita de bolo, onde você pode seguir os passos da receita para fazer o bolo,
// e depois pode usar a mesma receita para fazer outro bolo.

// As funções em GOLANG podem ter "parâmetros" e "return" (retornos)
// parâmetros são dados colocados na função para ela funcionar
// E return(retorno) é o que a função devolve para você


func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string { // declaração de função anônima, que é uma função sem nome, e atribui a variável f
		fmt.Println(txt) // imprime "Função f" no console
	}
	resultado := f("Texto da função 1") // chamada da função anônima, passando o argumento "Texto da função 1"
	fmt.Println(resultado)              // imprime o resultado da função anônima, que é "Texto da função 1"

}
