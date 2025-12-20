package main

import "fmt"

func main() {
	lista := []float64{98, 93, 77, 82, 83} // crie uma lista de numeros decimais com estes valores{}.
	total := 0.0 // crie uma variavel total começando em 0.
	for _, valor := range lista { // para cada valor dentro da lista
		total += valor // some o valor atual ao total.
	}
	fmt.Println(total / float64(len(lista))) // divindo a soma pela qtd(len) de notas.
}