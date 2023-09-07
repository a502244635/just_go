package main

import "fmt"

func main() {
	var i int = 10
	fmt.Printf("i = %v", i)

	var p *int = &i
	fmt.Printf(" 变量i地址 = %v", &i)

	fmt.Printf(" p地址是：%v", &p)

	fmt.Printf(" p指向的值是：%v", *p)
}
