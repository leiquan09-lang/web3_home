package main

import "fmt"

func addTen(p *int) {
	*p += 10
}
func main() {
	x := 5
	fmt.Println("调用前 x=", x)
	addTen(&x)
	fmt.Println("调用后 x=", x)
}
