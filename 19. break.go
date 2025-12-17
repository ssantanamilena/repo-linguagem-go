package main

import "fmt"

func main(){
	x := 0 
	for {if x < 20 {
		x++
		fmt.Println("x < 20")	
	}else{
		break
	}
}
}