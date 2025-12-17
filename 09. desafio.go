package main

import "fmt"

const ebulicaoK = 373.0

func main() {

	tempK := ebulicaoK

	tempC := tempK - 273.0

fmt.Println("A temperatura de ebulição da água em K:", tempK, "temperatura de ebulição da água em °C:", tempC)
}
