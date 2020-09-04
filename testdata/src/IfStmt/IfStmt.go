package IfStmt

func f() {
	var i int
	if true {
		i = i // want "variable is not initialized"
	}
}