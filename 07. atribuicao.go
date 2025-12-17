package main

import "fmt"

const ebulicaof float64 = 212.0 //declaração do ponto da variável

func main(){ //função principal
	var tempF float64 = ebulicaof
	var tempC float64 = (tempF - 32)*5/9 //conversão de F em C

	fmt.Println("A temperatura de ebulição da água em ºF = ", tempF)
	fmt.Println("A temperatura de ebulição da água em ºC = ", tempC)
	}
