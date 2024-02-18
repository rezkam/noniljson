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

func TestAnalyzerIgnoringGeneratedFiles(t *testing.T) {
	oldIgnoreGeneratedFiles := ignoreGeneratedFiles
	ignoreGeneratedFiles = true
	defer func() {
		ignoreGeneratedFiles = oldIgnoreGeneratedFiles
	}()
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "ignoregenerated")
}
