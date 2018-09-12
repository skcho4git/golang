//flag package
package main

import (
	
	"fmt"
	"flag"
)

func main() {
	
	var (
		
		max int
		message string
		x bool
	)
	
	
	flag.IntVar(&max,"n",32,"処理数の最大値")
	flag.StringVar(&message, "m", "", "処理メッセージ")
	flag.BoolVar(&x, "x", false, "拡張オプション")
	
	flag.Parse()
	
	fmt.Println("処理の最大値=",max)
	fmt.Println("処理メッセージ=",message)
	fmt.Println("拡張オプション=",x)
	
	
}