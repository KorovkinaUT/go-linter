package analyzer

import "golang.org/x/tools/go/analysis"

var Analyzer = &analysis.Analyzer{
	Name: "loglint",
	Doc:  "checks log messages formatting rules",
	Run:  run,
}
