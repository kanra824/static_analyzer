package ReturnStmt

func f () {
	i := g(1, 2)
	i = i
}

func g (a int, b int) int {
	return a + b
}