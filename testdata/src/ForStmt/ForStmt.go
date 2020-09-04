package ForStmt

func f() {
	var sz int
	for i := 0; i < sz; i++ { // want "variable is not initialized"
	}
	sz = 3

	var j int
	for i := j; i < sz; i++ { // want "variable is not initialized"
	}

	for i := 0; i < sz; i = i + j { // want "variable is not initialized"
	}


}
