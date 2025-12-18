package main

import "fmt"

func main(){
	fatia1 := []int{1,2,3}
	fatia2 := append(fatia1, 4, 5) // fatia da lista um + a lista 2.
	fmt.Println(fatia1, fatia2)
}