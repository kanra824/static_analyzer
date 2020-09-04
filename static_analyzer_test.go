package static_analyzer_test

import (
	"testing"

	"static_analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, static_analyzer.Analyzer, "a")
	analysistest.Run(t, testdata, static_analyzer.Analyzer, "b")
	analysistest.Run(t, testdata, static_analyzer.Analyzer, "c")
}

