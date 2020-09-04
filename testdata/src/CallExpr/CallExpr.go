package CallExpr

func f() {
	var i int
	i = g(i, 0) // want "variable is not initialized"
}

func g(a int, b int) int {
	return a + b
}