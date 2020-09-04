package IfStmt

func f() {
	var i int
	if true {
		i = i // want "variable is not initialized"
	}

	var j bool
	if j { // want "variable is not initialized"
		i = 1
	}

	var k bool
	if k := true; k {
		i = 1
	}
	k = k
}