package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/KorovkinaUT/go-linter/internal/analyzer"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
