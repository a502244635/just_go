package main

import "fmt"

func main() {
	var num int = 99
	var p *int = &num
	*p = 100
	fmt.Println(num)
}
