package main

import (
	"static_analyzer"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(static_analyzer.Analyzer) }

