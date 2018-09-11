package main

import (
	
	"fmt"
)


func integers() func() int {
	
	i := 0
	
	return func() int {
		
		i += 1
		return i
		
	}
	
	
	
}

func main() {
	
	f := integers()
	
	fmt.Println(f());
	fmt.Println(f());
	fmt.Println(f());	
	
}


