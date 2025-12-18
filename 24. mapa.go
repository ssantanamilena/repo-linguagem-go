package main

import "fmt"

func main(){
	x := make(map[string]int)
	x["chave"] = 10
	fmt.Println(x["chave"])
}