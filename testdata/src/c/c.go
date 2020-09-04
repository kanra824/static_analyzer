package c

func f() {
	var i int
	i = i + 0 // want "variable is not initialized"
}