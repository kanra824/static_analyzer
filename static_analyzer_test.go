package static_analyzer_test

import (
	"testing"

	"static_analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, static_analyzer.Analyzer, "Ident")
	analysistest.Run(t, testdata, static_analyzer.Analyzer, "BasicLit")
	analysistest.Run(t, testdata, static_analyzer.Analyzer, "BinaryExpr")
	analysistest.Run(t, testdata, static_analyzer.Analyzer, "CallExpr")
	analysistest.Run(t, testdata, static_analyzer.Analyzer, "ParenExpr")
	analysistest.Run(t, testdata, static_analyzer.Analyzer, "IfStmt")
}

