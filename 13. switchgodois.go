package main

import "fmt"

func main(){
	var opcao int
	fmt.Println("Digite uma opção")
	fmt.Println("1- Criar usuário")
	fmt.Println("2- Listar usuários")
	fmt.Println("3- Deletar usuário")

	fmt. Println("Opção: ")
	fmt.Scanln(&opcao) //& serve para o scanln saber onde salvar a resposta.

	switch opcao {
	case 1:
		fmt.Println("Criando usuário...")
	case 2:
		fmt.Println("Listando usuário...")	
	case 3:
		fmt.Println("Deletando usuário...")		
	default:
		fmt.Println("Opção inválida")
	}
}