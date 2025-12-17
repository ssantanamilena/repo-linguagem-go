package main

import "fmt"

func main(){
	for sala := 1; sala <= 3; sala++{
		fmt.Println("Sala", sala)

		for alunos :=1; alunos <= 4; alunos++{
			fmt.Println("Alunos", alunos)
		}
	}
}