package CallExpr

import (
	"fmt"
)

func f() {
	var i, j int
	fmt.Println(g(i, j))
}

func g(a int, b int) int {
	return a + b
}