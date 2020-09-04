package IncDecStmt

func f() {
	var i int
	i++ // want "variable is not initialized"
}