package main

import (
	
	"fmt"
)

func div(a,b int) (int,int){
	q := a / b
	r := a % b
	return q,r
	
	
}
func main() {
	
	fmt.Println("Hello World!!")
	
	var x, y int
	x, y = 1, 2
	
	fmt.Println(x,y)
	
	q,r := div(19,7)
	fmt.Printf("商=%d 余剰=%d\n",q,r)
	
	f := func(x, y int) int { return x + y }
	fmt.Println(f(1,2))
	
}


