package main

import "fmt"

func main(){
	for x := 0; x <20; x++{
		if x == 3 { 
			continue // pula o número 3
		}
		fmt.Println(x)
	}

}