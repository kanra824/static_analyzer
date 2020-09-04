package Ident

func f() {
	var i int
	i = i // want "variable is not initialized"
}

