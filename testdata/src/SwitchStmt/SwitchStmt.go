package SwitchStmt

func f() {
	var i int
	switch i { // want "variable is not initialized"
	case 0:
	case 1:
	default:
	}

	switch {
	case i == 0: // want "variable is not initialized"
	case i == 1: // want "variable is not initialized"
	default:
	}

	switch {
	default:
		i = i // want "variable is not initialized"
	}

	i = i // want "variable is not initialized"
}
