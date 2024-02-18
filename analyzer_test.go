package noniljson

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	for _, pkg := range []string{"pointers", "slices", "maps", "interfaces", "customtypes", "withomitempty", "nonnullables"} {
		analysistest.Run(t, testdata, Analyzer, pkg)
	}
}
