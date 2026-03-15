package analyzer

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzer = &analysis.Analyzer{
	Name:     "loglint",
	Doc:      "checks log messages formatting rules",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}
