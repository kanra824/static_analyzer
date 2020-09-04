package SelectorExpr

import "fmt"

type A struct {
	a int
	b int
}

func f() {
	var a A
	i := a.a // want "variable is not initialized"
	fmt.Println(i)
}